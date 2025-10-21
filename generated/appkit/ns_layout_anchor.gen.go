// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutAnchor] class.
var (
	LayoutAnchorClass     _LayoutAnchorClass
	LayoutAnchorClassOnce sync.Once
)

func getLayoutAnchorClass() _LayoutAnchorClass {
	LayoutAnchorClassOnce.Do(func() {
		LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}
	})
	return LayoutAnchorClass
}

type _LayoutAnchorClass struct {
	class objc.Class
}

// An interface definition for the [LayoutAnchor] class.
type ILayoutAnchor interface {
	objectivec.IObject
}

// A factory class for creating layout constraint objects using a fluent API.
//
// Use these constraints to programatically define your layout using Auto Layout. Instead of creating objects directly, start with an or object you wish to constrain, and select one of that object’s anchor properties. These properties correspond to the main values used in Auto Layout, and provide an appropriate subclass for creating constraints to that attribute. Use the anchor’s methods to construct your constraint. As you can see from these examples, the class provides several advantages over using the API directly. The code is cleaner, more concise, and easier to read. The subclasses provide additional type checking, preventing you from creating invalid constraints. For more information on the anchor properties, see in the or .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor
type LayoutAnchor struct {
	objectivec.Object
}

// LayoutAnchorFrom constructs a [LayoutAnchor] from an unsafe.Pointer.
//
// A factory class for creating layout constraint objects using a fluent API.
func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutAnchor) Autorelease() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutAnchor creates a new LayoutAnchor instance.
func NewLayoutAnchor() LayoutAnchor {
	return getLayoutAnchorClass().New()
}


// A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/hasAmbiguousLayout
func (l_ LayoutAnchor) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}

// The constraints that impact the layout of the anchor.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/constraintsaffectinglayout
func (l_ LayoutAnchor) ConstraintsAffectingLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("constraintsAffectingLayout"))
	return rv
}


// SetConstraintsAffectingLayout sets the value of the constraintsAffectingLayout property.
// The constraints that impact the layout of the anchor.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/constraintsaffectinglayout
func (l_ LayoutAnchor) SetConstraintsAffectingLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setConstraintsAffectingLayout:"), value)
}

// The layout item used to calculate the anchor’s position.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/item
func (l_ LayoutAnchor) Item() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("item"))
	return rv
}


// SetItem sets the value of the item property.
// The layout item used to calculate the anchor’s position.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/item
func (l_ LayoutAnchor) SetItem(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setItem:"), value)
}

// The name assigned to the anchor for debugging purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/name
func (l_ LayoutAnchor) Name() string {
	rv := objc.Send[string](l_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name assigned to the anchor for debugging purposes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutanchor/name
func (l_ LayoutAnchor) SetName(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), objc.String(value))
}

// A layout anchor representing the bottom edge of the view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (l_ LayoutAnchor) BottomAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("bottomAnchor"))
	return rv
}


// SetBottomAnchor sets the value of the bottomAnchor property.
// A layout anchor representing the bottom edge of the view’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (l_ LayoutAnchor) SetBottomAnchor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBottomAnchor:"), value)
}

// A layout anchor representing the leading edge of the view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (l_ LayoutAnchor) LeadingAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("leadingAnchor"))
	return rv
}


// SetLeadingAnchor sets the value of the leadingAnchor property.
// A layout anchor representing the leading edge of the view’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (l_ LayoutAnchor) SetLeadingAnchor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeadingAnchor:"), value)
}

// A layout anchor representing the left edge of the view’s frame.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (l_ LayoutAnchor) LeftAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("leftAnchor"))
	return rv
}


// SetLeftAnchor sets the value of the leftAnchor property.
// A layout anchor representing the left edge of the view’s frame.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (l_ LayoutAnchor) SetLeftAnchor(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeftAnchor:"), value)
}



