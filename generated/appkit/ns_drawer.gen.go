// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// A user interface element that contains and displays text, scroll, and browser views, in addition to other view subclasses.
//
// A drawer is associated with a window, called its parent, and can appear only while its parent is visible onscreen. A drawer cannot be moved or ordered independently of a window, but is instead attached to one edge of its parent and moves along with it.
//
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
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentsize
func (d_ Drawer) ContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](d_.ID, objc.Sel("contentSize"))
	return rv
}


// SetContentSize sets the value of the contentSize property.
// The size of the receiver’s content area.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentsize
func (d_ Drawer) SetContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentSize:"), value)
}

// The receiver’s content view.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentview
func (d_ Drawer) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The receiver’s content view.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/contentview
func (d_ Drawer) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentView:"), value)
}

// The receiver’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/delegate
func (d_ Drawer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The receiver’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/delegate
func (d_ Drawer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}

// The edge of the window that the receiver is connected to.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/edge
func (d_ Drawer) Edge() int {
	rv := objc.Send[int](d_.ID, objc.Sel("edge"))
	return rv
}


// SetEdge sets the value of the edge property.
// The edge of the window that the receiver is connected to.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/edge
func (d_ Drawer) SetEdge(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEdge:"), value)
}

// The receiver’s leading offset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/leadingoffset
func (d_ Drawer) LeadingOffset() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("leadingOffset"))
	return rv
}


// SetLeadingOffset sets the value of the leadingOffset property.
// The receiver’s leading offset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/leadingoffset
func (d_ Drawer) SetLeadingOffset(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLeadingOffset:"), value)
}

// The maximum allowed size of the receiver’s content area.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/maxcontentsize
func (d_ Drawer) MaxContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](d_.ID, objc.Sel("maxContentSize"))
	return rv
}


// SetMaxContentSize sets the value of the maxContentSize property.
// The maximum allowed size of the receiver’s content area.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/maxcontentsize
func (d_ Drawer) SetMaxContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxContentSize:"), value)
}

// The minimum allowed size of the receiver’s content area.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/mincontentsize
func (d_ Drawer) MinContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](d_.ID, objc.Sel("minContentSize"))
	return rv
}


// SetMinContentSize sets the value of the minContentSize property.
// The minimum allowed size of the receiver’s content area.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/mincontentsize
func (d_ Drawer) SetMinContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinContentSize:"), value)
}

// The receiver’s parent window.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/parentwindow
func (d_ Drawer) ParentWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("parentWindow"))
	return rv
}


// SetParentWindow sets the value of the parentWindow property.
// The receiver’s parent window.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/parentwindow
func (d_ Drawer) SetParentWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParentWindow:"), value)
}

// The receiver’s preferred, or default, edge.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/preferrededge
func (d_ Drawer) PreferredEdge() int {
	rv := objc.Send[int](d_.ID, objc.Sel("preferredEdge"))
	return rv
}


// SetPreferredEdge sets the value of the preferredEdge property.
// The receiver’s preferred, or default, edge.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/preferrededge
func (d_ Drawer) SetPreferredEdge(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredEdge:"), value)
}

// The state of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/state-swift.property
func (d_ Drawer) State() int {
	rv := objc.Send[int](d_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The state of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/state-swift.property
func (d_ Drawer) SetState(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setState:"), value)
}

// The receiver’s trailing offset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/trailingoffset
func (d_ Drawer) TrailingOffset() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("trailingOffset"))
	return rv
}


// SetTrailingOffset sets the value of the trailingOffset property.
// The receiver’s trailing offset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdrawer/trailingoffset
func (d_ Drawer) SetTrailingOffset(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTrailingOffset:"), value)
}



