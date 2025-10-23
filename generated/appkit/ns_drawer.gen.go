// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Drawer] class.
var (
	DrawerClass     _DrawerClass
	DrawerClassOnce sync.Once
)

func getDrawerClass() _DrawerClass {
	DrawerClassOnce.Do(func() {
		DrawerClass = _DrawerClass{objc.GetClass("NSDrawer")}
	})
	return DrawerClass
}

type _DrawerClass struct {
	class objc.Class
}

// An interface definition for the [Drawer] class.
type IDrawer interface {
	IResponder
	// properties:
	ContentSize() objc.IObject /* cross-framework: Size */
	SetContentSize(value objc.IObject /* cross-framework: Size */)
	ContentView() IView
	SetContentView(value IView)
	Delegate() DrawerDelegate /* not a class type */
	SetDelegate(value DrawerDelegate /* not a class type */)
	Edge() RectEdge /* not a class type */
	SetEdge(value RectEdge /* not a class type */)
	LeadingOffset() float64 /* primitive/slice/pointer. */
	SetLeadingOffset(value float64 /* primitive/slice/pointer. */)
	MaxContentSize() objc.IObject /* cross-framework: Size */
	SetMaxContentSize(value objc.IObject /* cross-framework: Size */)
	MinContentSize() objc.IObject /* cross-framework: Size */
	SetMinContentSize(value objc.IObject /* cross-framework: Size */)
	ParentWindow() IWindow
	SetParentWindow(value IWindow)
	PreferredEdge() RectEdge /* not a class type */
	SetPreferredEdge(value RectEdge /* not a class type */)
	State() int /* primitive/slice/pointer. */
	SetState(value int /* primitive/slice/pointer. */)
	TrailingOffset() float64 /* primitive/slice/pointer. */
	SetTrailingOffset(value float64 /* primitive/slice/pointer. */)
	// methods:
}

// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses.
//
// A drawer is associated with a window, called its parent, and can appear only while its parent is visible onscreen. A drawer cannot be moved or ordered independently of a window, but is instead attached to one edge of its parent and moves along with it.


// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer
type Drawer struct {
	Responder
}

// DrawerFrom constructs a [Drawer] from an unsafe.Pointer.
//
// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses.
func DrawerFrom(ptr unsafe.Pointer) Drawer {
	return Drawer{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DrawerClass) New() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Drawer) Init() Drawer {
	rv := objc.Send[Drawer](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Drawer) Autorelease() Drawer {
	rv := objc.Send[Drawer](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDrawer creates a new Drawer instance.
func NewDrawer() Drawer {
	return getDrawerClass().New()
}



// The size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentsize
func (d_ Drawer) ContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("contentSize"))
	return rv
}


// The size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentsize
func (d_ Drawer) SetContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentSize:"), value)
}


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentview
func (d_ Drawer) ContentView() IView {
	rv := objc.Send[View](d_.ID, objc.Sel("contentView"))
	return rv
}


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentview
func (d_ Drawer) SetContentView(value IView) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentView:"), value)
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/delegate
func (d_ Drawer) Delegate() DrawerDelegate /* not a class type */ {
	rv := objc.Send[DrawerDelegate](d_.ID, objc.Sel("delegate"))
	return rv
}


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/delegate
func (d_ Drawer) SetDelegate(value DrawerDelegate /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}


// The edge of the window that the receiver is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/edge
func (d_ Drawer) Edge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](d_.ID, objc.Sel("edge"))
	return rv
}


// The edge of the window that the receiver is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/edge
func (d_ Drawer) SetEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEdge:"), value)
}


// The receiver’s leading offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/leadingoffset
func (d_ Drawer) LeadingOffset() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("leadingOffset"))
	return rv
}


// The receiver’s leading offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/leadingoffset
func (d_ Drawer) SetLeadingOffset(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLeadingOffset:"), value)
}


// The maximum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/maxcontentsize
func (d_ Drawer) MaxContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("maxContentSize"))
	return rv
}


// The maximum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/maxcontentsize
func (d_ Drawer) SetMaxContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxContentSize:"), value)
}


// The minimum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/mincontentsize
func (d_ Drawer) MinContentSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("minContentSize"))
	return rv
}


// The minimum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/mincontentsize
func (d_ Drawer) SetMinContentSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinContentSize:"), value)
}


// The receiver’s parent window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/parentwindow
func (d_ Drawer) ParentWindow() IWindow {
	rv := objc.Send[Window](d_.ID, objc.Sel("parentWindow"))
	return rv
}


// The receiver’s parent window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/parentwindow
func (d_ Drawer) SetParentWindow(value IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParentWindow:"), value)
}


// The receiver’s preferred, or default, edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/preferrededge
func (d_ Drawer) PreferredEdge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](d_.ID, objc.Sel("preferredEdge"))
	return rv
}


// The receiver’s preferred, or default, edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/preferrededge
func (d_ Drawer) SetPreferredEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredEdge:"), value)
}


// The state of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/state-swift.property
func (d_ Drawer) State() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](d_.ID, objc.Sel("state"))
	return rv
}


// The state of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/state-swift.property
func (d_ Drawer) SetState(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setState:"), value)
}


// The receiver’s trailing offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/trailingoffset
func (d_ Drawer) TrailingOffset() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("trailingOffset"))
	return rv
}


// The receiver’s trailing offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/trailingoffset
func (d_ Drawer) SetTrailingOffset(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTrailingOffset:"), value)
}



