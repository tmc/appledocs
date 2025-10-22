// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [LayoutGuide] class.
type ILayoutGuide interface {
	objectivec.IObject
	BottomAnchor() NSLayoutYAxisAnchor
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value IUserInterfaceItemIdentifier)
	LeadingAnchor() NSLayoutXAxisAnchor
	CenterXAnchor() NSLayoutXAxisAnchor
	SetCenterXAnchor(value ILayoutXAxisAnchor)
	CenterYAnchor() NSLayoutYAxisAnchor
	SetCenterYAnchor(value ILayoutYAxisAnchor)
	Frame() coregraphics.CGRect
	SetFrame(value coregraphics.CGRect)
	HasAmbiguousLayout() bool
	SetHasAmbiguousLayout(value bool)
	HeightAnchor() NSLayoutDimension
	SetHeightAnchor(value ILayoutDimension)
	LeftAnchor() NSLayoutXAxisAnchor
	SetLeftAnchor(value ILayoutXAxisAnchor)
	OwningView() NSView
	SetOwningView(value IView)
	RightAnchor() NSLayoutXAxisAnchor
	SetRightAnchor(value ILayoutXAxisAnchor)
	TopAnchor() NSLayoutYAxisAnchor
	SetTopAnchor(value ILayoutYAxisAnchor)
	TrailingAnchor() NSLayoutXAxisAnchor
	SetTrailingAnchor(value ILayoutXAxisAnchor)
	WidthAnchor() NSLayoutDimension
	SetWidthAnchor(value ILayoutDimension)
}

// A rectangular area that can interact with Auto Layout.
//
// Use layout guides to replace the placeholder views you may have created to represent inter-view spaces or encapsulation in your user interface. Traditionally, there were a number of Auto Layout techniques that required placeholder views. A placeholder view is an empty view that does not have any visual elements of its own and serves only to define a rectangular region in the view hierarchy. For example, if you wanted to use constraints to define the size or location of an empty space between views, you needed to use a placeholder view to represent that space. If you wanted to center a group of objects, you needed a placeholder view to contain those objects. Similarly, placeholder views could be used to contain and encapsulate part of your user interface. Placeholder views let you break up a large, complex user interface into self-contained, modular chunks. When used properly, they could greatly simplify your Auto Layout constraint logic. There are a number of costs associated with adding placeholder views to your view hierarchy. First, there is the cost of creating and maintaining the view itself. Second, the placeholder view is a full member of the view hierarchy, which means that it adds overhead to every task the hierarchy performs. Worst of all, the invisible placeholder view can intercept messages that are intended for other views, causing problems that are very difficult to find. The class is designed to perform all the tasks previously performed by placeholder views, but to do it in a safer, more efficient manner. Layout guides are not views. They do not use as much memory, and they do not participate in the view hierarchy. Instead, they simply define a rectangular region in their owning view’s coordinate system that can interact with Auto Layout.
//
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

// Alloc allocates a new instance without initialization.
func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A layout anchor representing the bottom edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/bottomAnchor
func (l_ LayoutGuide) BottomAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](l_.ID, objc.Sel("bottomAnchor"))
	return rv
}

// A string used to identify the layout guide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := objc.Send[UserInterfaceItemIdentifier](l_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A string used to identify the layout guide.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) SetIdentifier(value IUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}

// A layout anchor representing the leading edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leadingAnchor
func (l_ LayoutGuide) LeadingAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l_.ID, objc.Sel("leadingAnchor"))
	return rv
}

// A layout anchor representing the horizontal center of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/centerxanchor
func (l_ LayoutGuide) CenterXAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l_.ID, objc.Sel("centerXAnchor"))
	return rv
}


// SetCenterXAnchor sets the value of the centerXAnchor property.
// A layout anchor representing the horizontal center of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/centerxanchor
func (l_ LayoutGuide) SetCenterXAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCenterXAnchor:"), value)
}

// A layout anchor representing the vertical center of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/centeryanchor
func (l_ LayoutGuide) CenterYAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](l_.ID, objc.Sel("centerYAnchor"))
	return rv
}


// SetCenterYAnchor sets the value of the centerYAnchor property.
// A layout anchor representing the vertical center of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/centeryanchor
func (l_ LayoutGuide) SetCenterYAnchor(value ILayoutYAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCenterYAnchor:"), value)
}

// The layout guide’s frame in its owning view’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/frame
func (l_ LayoutGuide) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("frame"))
	return rv
}


// SetFrame sets the value of the frame property.
// The layout guide’s frame in its owning view’s coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/frame
func (l_ LayoutGuide) SetFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFrame:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/hasambiguouslayout
func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}


// SetHasAmbiguousLayout sets the value of the hasAmbiguousLayout property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/hasambiguouslayout
func (l_ LayoutGuide) SetHasAmbiguousLayout(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHasAmbiguousLayout:"), value)
}

// A layout anchor representing the height of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/heightanchor
func (l_ LayoutGuide) HeightAnchor() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](l_.ID, objc.Sel("heightAnchor"))
	return rv
}


// SetHeightAnchor sets the value of the heightAnchor property.
// A layout anchor representing the height of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/heightanchor
func (l_ LayoutGuide) SetHeightAnchor(value ILayoutDimension) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHeightAnchor:"), value)
}

// A layout anchor representing the left edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/leftanchor
func (l_ LayoutGuide) LeftAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l_.ID, objc.Sel("leftAnchor"))
	return rv
}


// SetLeftAnchor sets the value of the leftAnchor property.
// A layout anchor representing the left edge of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/leftanchor
func (l_ LayoutGuide) SetLeftAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeftAnchor:"), value)
}

// The view that owns this layout guide.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/owningview
func (l_ LayoutGuide) OwningView() NSView {
	rv := objc.Send[NSView](l_.ID, objc.Sel("owningView"))
	return rv
}


// SetOwningView sets the value of the owningView property.
// The view that owns this layout guide.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/owningview
func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOwningView:"), value)
}

// A layout anchor representing the right edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/rightanchor
func (l_ LayoutGuide) RightAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l_.ID, objc.Sel("rightAnchor"))
	return rv
}


// SetRightAnchor sets the value of the rightAnchor property.
// A layout anchor representing the right edge of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/rightanchor
func (l_ LayoutGuide) SetRightAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRightAnchor:"), value)
}

// A layout anchor representing the top edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/topanchor
func (l_ LayoutGuide) TopAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](l_.ID, objc.Sel("topAnchor"))
	return rv
}


// SetTopAnchor sets the value of the topAnchor property.
// A layout anchor representing the top edge of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/topanchor
func (l_ LayoutGuide) SetTopAnchor(value ILayoutYAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTopAnchor:"), value)
}

// A layout anchor representing the trailing edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/trailinganchor
func (l_ LayoutGuide) TrailingAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](l_.ID, objc.Sel("trailingAnchor"))
	return rv
}


// SetTrailingAnchor sets the value of the trailingAnchor property.
// A layout anchor representing the trailing edge of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/trailinganchor
func (l_ LayoutGuide) SetTrailingAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTrailingAnchor:"), value)
}

// A layout anchor representing the width of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/widthanchor
func (l_ LayoutGuide) WidthAnchor() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](l_.ID, objc.Sel("widthAnchor"))
	return rv
}


// SetWidthAnchor sets the value of the widthAnchor property.
// A layout anchor representing the width of the layout guide’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/widthanchor
func (l_ LayoutGuide) SetWidthAnchor(value ILayoutDimension) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWidthAnchor:"), value)
}



