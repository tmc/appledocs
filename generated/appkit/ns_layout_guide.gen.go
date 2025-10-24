// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLayoutGuide */


/* debug [class_header]: Header for NSLayoutGuide */
// The class instance for the [LayoutGuide] class.
var (
	LayoutGuideClass     _LayoutGuideClass
	LayoutGuideClassOnce sync.Once
)

func getLayoutGuideClass() _LayoutGuideClass {
	LayoutGuideClassOnce.Do(func() {
		LayoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}
	})
	return LayoutGuideClass
}

type _LayoutGuideClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LayoutGuide */
// An interface definition for the [LayoutGuide] class.
type ILayoutGuide interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LayoutGuide */
	// properties:
	BottomAnchor() ILayoutYAxisAnchor
	CenterXAnchor() ILayoutXAxisAnchor
	CenterYAnchor() ILayoutYAxisAnchor
	Frame() Rect /* not a class type */
	HasAmbiguousLayout() bool
	HeightAnchor() ILayoutDimension
	Identifier() UserInterfaceItemIdentifier /* typedef */
	SetIdentifier(value UserInterfaceItemIdentifier /* typedef */)
	LeadingAnchor() ILayoutXAxisAnchor
	LeftAnchor() ILayoutXAxisAnchor
	OwningView() IView
	SetOwningView(value IView)
	RightAnchor() ILayoutXAxisAnchor
	TopAnchor() ILayoutYAxisAnchor
	TrailingAnchor() ILayoutXAxisAnchor
	WidthAnchor() ILayoutDimension
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LayoutGuide */
	// methods:
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LayoutGuide */
// Alloc allocates a new instance without initialization.
func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.Send[LayoutGuide](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutGuide) Autorelease() LayoutGuide {
	rv := objc.Send[LayoutGuide](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutGuide creates a new LayoutGuide instance.
func NewLayoutGuide() LayoutGuide {
	return getLayoutGuideClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LayoutGuide */
// A rectangular area that can interact with Auto Layout.
//
// Use layout guides to replace the placeholder views you may have created to represent inter-view spaces or encapsulation in your user interface. Traditionally, there were a number of Auto Layout techniques that required placeholder views. A placeholder view is an empty view that does not have any visual elements of its own and serves only to define a rectangular region in the view hierarchy. For example, if you wanted to use constraints to define the size or location of an empty space between views, you needed to use a placeholder view to represent that space. If you wanted to center a group of objects, you needed a placeholder view to contain those objects. Similarly, placeholder views could be used to contain and encapsulate part of your user interface. Placeholder views let you break up a large, complex user interface into self-contained, modular chunks. When used properly, they could greatly simplify your Auto Layout constraint logic. There are a number of costs associated with adding placeholder views to your view hierarchy. First, there is the cost of creating and maintaining the view itself. Second, the placeholder view is a full member of the view hierarchy, which means that it adds overhead to every task the hierarchy performs. Worst of all, the invisible placeholder view can intercept messages that are intended for other views, causing problems that are very difficult to find. The class is designed to perform all the tasks previously performed by placeholder views, but to do it in a safer, more efficient manner. Layout guides are not views. They do not use as much memory, and they do not participate in the view hierarchy. Instead, they simply define a rectangular region in their owning view’s coordinate system that can interact with Auto Layout.


// A rectangular area that can interact with Auto Layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide
type LayoutGuide struct {
	objectivec.Object
}

// LayoutGuideFrom constructs a [LayoutGuide] from an unsafe.Pointer.
//
// A rectangular area that can interact with Auto Layout.
func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LayoutGuide *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LayoutGuide */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LayoutGuide */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LayoutGuide */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/constraintsAffectingLayout(for:)
func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](l_.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ConstraintsAffectingLayoutForOrientation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LayoutGuide */

// A layout anchor representing the bottom edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/bottomAnchor
func (l_ LayoutGuide) BottomAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("bottomAnchor"))
	return rv
}/* debug [instance_properties/getter]: bottomAnchor */


// A layout anchor representing the horizontal center of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/centerXAnchor
func (l_ LayoutGuide) CenterXAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("centerXAnchor"))
	return rv
}/* debug [instance_properties/getter]: centerXAnchor */


// A layout anchor representing the vertical center of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/centerYAnchor
func (l_ LayoutGuide) CenterYAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("centerYAnchor"))
	return rv
}/* debug [instance_properties/getter]: centerYAnchor */


// The layout guide’s frame in its owning view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/frame
func (l_ LayoutGuide) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](l_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/hasAmbiguousLayout
func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}/* debug [instance_properties/getter]: hasAmbiguousLayout */


// A layout anchor representing the height of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/heightAnchor
func (l_ LayoutGuide) HeightAnchor() ILayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("heightAnchor"))
	return rv
}/* debug [instance_properties/getter]: heightAnchor */


// A string used to identify the layout guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A string used to identify the layout guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A layout anchor representing the leading edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leadingAnchor
func (l_ LayoutGuide) LeadingAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("leadingAnchor"))
	return rv
}/* debug [instance_properties/getter]: leadingAnchor */


// A layout anchor representing the left edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leftAnchor
func (l_ LayoutGuide) LeftAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("leftAnchor"))
	return rv
}/* debug [instance_properties/getter]: leftAnchor */


// The view that owns this layout guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/owningView
func (l_ LayoutGuide) OwningView() IView {
	rv := objc.Send[View](l_.ID, objc.Sel("owningView"))
	return rv
}/* debug [instance_properties/getter]: owningView */


// The view that owns this layout guide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/owningView
func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOwningView:"), value)
}/* debug [instance_properties/setter]: owningView */


// A layout anchor representing the right edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/rightAnchor
func (l_ LayoutGuide) RightAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("rightAnchor"))
	return rv
}/* debug [instance_properties/getter]: rightAnchor */


// A layout anchor representing the top edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/topAnchor
func (l_ LayoutGuide) TopAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("topAnchor"))
	return rv
}/* debug [instance_properties/getter]: topAnchor */


// A layout anchor representing the trailing edge of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/trailingAnchor
func (l_ LayoutGuide) TrailingAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("trailingAnchor"))
	return rv
}/* debug [instance_properties/getter]: trailingAnchor */


// A layout anchor representing the width of the layout guide’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/widthAnchor
func (l_ LayoutGuide) WidthAnchor() ILayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("widthAnchor"))
	return rv
}/* debug [instance_properties/getter]: widthAnchor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLayoutGuide */



