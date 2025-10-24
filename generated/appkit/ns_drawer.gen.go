// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDrawer */


/* debug [class_header]: Header for NSDrawer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Drawer */
// An interface definition for the [Drawer] class.
type IDrawer interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for Drawer */
	// properties:
	ContentSize() Size /* not a class type */
	SetContentSize(value Size /* not a class type */)
	ContentView() IView
	SetContentView(value IView)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Edge() RectEdge /* not a class type */
	LeadingOffset() float64
	SetLeadingOffset(value float64)
	MaxContentSize() Size /* not a class type */
	SetMaxContentSize(value Size /* not a class type */)
	MinContentSize() Size /* not a class type */
	SetMinContentSize(value Size /* not a class type */)
	ParentWindow() IWindow
	SetParentWindow(value IWindow)
	PreferredEdge() RectEdge /* not a class type */
	SetPreferredEdge(value RectEdge /* not a class type */)
	State() int
	TrailingOffset() float64
	SetTrailingOffset(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Drawer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Drawer */
// Alloc allocates a new instance without initialization.
func (dc _DrawerClass) Alloc() Drawer {
	rv := objc.Send[Drawer](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Drawer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Drawer */

// Creates a new drawer with the given size on the specified edge of the parent window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/init(contentSize:preferredEdge:)
func NewDrawerWithContentSizePreferredEdge(contentSize Size /* not a class type */, edge RectEdge /* not a class type */) Drawer {
	instance := getDrawerClass().Alloc()
	rv := objc.Send[Drawer](instance.ID, objc.Sel("initWithContentSize:preferredEdge:"), contentSize, edge)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDrawerWithContentSizePreferredEdge */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Drawer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Drawer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Drawer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Drawer */

// The size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/contentSize
func (d_ Drawer) ContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("contentSize"))
	return rv
}/* debug [instance_properties/getter]: contentSize */


// The size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/contentSize
func (d_ Drawer) SetContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentSize:"), value)
}/* debug [instance_properties/setter]: contentSize */


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/contentView
func (d_ Drawer) ContentView() IView {
	rv := objc.Send[View](d_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/contentView
func (d_ Drawer) SetContentView(value IView) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/delegate
func (d_ Drawer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/delegate
func (d_ Drawer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The edge of the window that the receiver is connected to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/edge
func (d_ Drawer) Edge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](d_.ID, objc.Sel("edge"))
	return rv
}/* debug [instance_properties/getter]: edge */


// The receiver’s leading offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/leadingOffset
func (d_ Drawer) LeadingOffset() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("leadingOffset"))
	return rv
}/* debug [instance_properties/getter]: leadingOffset */


// The receiver’s leading offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/leadingOffset
func (d_ Drawer) SetLeadingOffset(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLeadingOffset:"), value)
}/* debug [instance_properties/setter]: leadingOffset */


// The maximum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/maxContentSize
func (d_ Drawer) MaxContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("maxContentSize"))
	return rv
}/* debug [instance_properties/getter]: maxContentSize */


// The maximum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/maxContentSize
func (d_ Drawer) SetMaxContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxContentSize:"), value)
}/* debug [instance_properties/setter]: maxContentSize */


// The minimum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/minContentSize
func (d_ Drawer) MinContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](d_.ID, objc.Sel("minContentSize"))
	return rv
}/* debug [instance_properties/getter]: minContentSize */


// The minimum allowed size of the receiver’s content area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/minContentSize
func (d_ Drawer) SetMinContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinContentSize:"), value)
}/* debug [instance_properties/setter]: minContentSize */


// The receiver’s parent window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/parentWindow
func (d_ Drawer) ParentWindow() IWindow {
	rv := objc.Send[Window](d_.ID, objc.Sel("parentWindow"))
	return rv
}/* debug [instance_properties/getter]: parentWindow */


// The receiver’s parent window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/parentWindow
func (d_ Drawer) SetParentWindow(value IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParentWindow:"), value)
}/* debug [instance_properties/setter]: parentWindow */


// The receiver’s preferred, or default, edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/preferredEdge
func (d_ Drawer) PreferredEdge() RectEdge /* not a class type */ {
	rv := objc.Send[RectEdge](d_.ID, objc.Sel("preferredEdge"))
	return rv
}/* debug [instance_properties/getter]: preferredEdge */


// The receiver’s preferred, or default, edge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/preferredEdge
func (d_ Drawer) SetPreferredEdge(value RectEdge /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredEdge:"), value)
}/* debug [instance_properties/setter]: preferredEdge */


// The state of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/state-swift.property
func (d_ Drawer) State() int {
	rv := objc.Send[int](d_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The receiver’s trailing offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/trailingOffset
func (d_ Drawer) TrailingOffset() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("trailingOffset"))
	return rv
}/* debug [instance_properties/getter]: trailingOffset */


// The receiver’s trailing offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/trailingOffset
func (d_ Drawer) SetTrailingOffset(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTrailingOffset:"), value)
}/* debug [instance_properties/setter]: trailingOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDrawer */


