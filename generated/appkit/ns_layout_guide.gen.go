// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (l_ LayoutGuide) BottomAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("bottomAnchor"))
	return rv
}

// A string used to identify the layout guide.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A string used to identify the layout guide.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l_ LayoutGuide) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIdentifier:"), value)
}

// A layout anchor representing the leading edge of the layout guide’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leadingAnchor
func (l_ LayoutGuide) LeadingAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("leadingAnchor"))
	return rv
}



