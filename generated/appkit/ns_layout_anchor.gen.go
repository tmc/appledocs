// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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



