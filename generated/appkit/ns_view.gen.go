// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [View] class.
var (
	ViewClass     _ViewClass
	ViewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	ViewClassOnce.Do(func() {
		ViewClass = _ViewClass{objc.GetClass("NSView")}
	})
	return ViewClass
}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	IResponder
	// properties:
	Bounds() objc.IObject /* cross-framework: Rect */
	SetBounds(value objc.IObject /* cross-framework: Rect */)
	Frame() objc.IObject /* cross-framework: Rect */
	SetFrame(value objc.IObject /* cross-framework: Rect */)
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	PrefersCompactControlSizeMetrics() bool
	SetPrefersCompactControlSizeMetrics(value bool)
	Window() objc.IObject /* cross-framework: Window */
	SetWindow(value objc.IObject /* cross-framework: Window */)
	WritingToolsCoordinator() IWritingToolsCoordinator
	SetWritingToolsCoordinator(value IWritingToolsCoordinator)
	// methods:
}

// The infrastructure for drawing, printing, and handling events in an app.
//
// You typically don’t use objects directly. Instead, you use objects that descend from or you subclass yourself and override its methods to implement the behavior you need. An instance of the class (or one of its subclasses) is commonly known as a view object, or simply as a view. Views handle the presentation and interaction with your app’s visible content. You arrange one or more views inside an object, which acts as a wrapper for your content. A view object defines a rectangular region for drawing and receiving mouse events. Views handle other chores as well, including the dragging of icons and working with the class to support efficient scrolling. AppKit handles most of your app’s management. Unless you’re implementing a concrete subclass of or working intimately with the content of the view hierarchy at runtime, you don’t need to know much about this class’s interface. For any view, there are many methods that you can use as-is. The following methods are commonly used. returns the location and size of the object. returns the internal origin and size of the object. determines whether the object needs to be redrawn. returns the object that contains the object. draws the object. (All subclasses must implement this method, but it’s rarely invoked explicitly.) An alternative to drawing is to update the layer directly using the method. For more information on how instances handle event and action messages, see . For more information on displaying tooltips and contextual menus, see and .


// The infrastructure for drawing, printing, and handling events in an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView
type View struct {
	Responder
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// The infrastructure for drawing, printing, and handling events in an app.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return getViewClass().New()
}



// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bounds
func (v_ View) Bounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](v_.ID, objc.Sel("bounds"))
	return rv
}


// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bounds
func (v_ View) SetBounds(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBounds:"), value)
}


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/frame
func (v_ View) Frame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](v_.ID, objc.Sel("frame"))
	return rv
}


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/frame
func (v_ View) SetFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrame:"), value)
}


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/needsdisplay
func (v_ View) NeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsDisplay"))
	return rv
}


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/needsdisplay
func (v_ View) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplay:"), value)
}


// When this property is true, any NSControls in the view or its descendants will be sized with compact
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/preferscompactcontrolsizemetrics
func (v_ View) PrefersCompactControlSizeMetrics() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("prefersCompactControlSizeMetrics"))
	return rv
}


// When this property is true, any NSControls in the view or its descendants will be sized with compact
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/preferscompactcontrolsizemetrics
func (v_ View) SetPrefersCompactControlSizeMetrics(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPrefersCompactControlSizeMetrics:"), value)
}


// The view’s window object, if it is installed in a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/window
func (v_ View) Window() objc.IObject /* cross-framework: Window */ {
	rv := objc.Send[Window](v_.ID, objc.Sel("window"))
	return rv
}


// The view’s window object, if it is installed in a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/window
func (v_ View) SetWindow(value objc.IObject /* cross-framework: Window */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWindow:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (v_ View) WritingToolsCoordinator() IWritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](v_.ID, objc.Sel("writingToolsCoordinator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/writingtoolscoordinator
func (v_ View) SetWritingToolsCoordinator(value IWritingToolsCoordinator) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}



