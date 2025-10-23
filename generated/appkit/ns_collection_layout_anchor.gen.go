// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionLayoutAnchor] class.
var (
	CollectionLayoutAnchorClass     _CollectionLayoutAnchorClass
	CollectionLayoutAnchorClassOnce sync.Once
)

func getCollectionLayoutAnchorClass() _CollectionLayoutAnchorClass {
	CollectionLayoutAnchorClassOnce.Do(func() {
		CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}
	})
	return CollectionLayoutAnchorClass
}

type _CollectionLayoutAnchorClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutAnchor] class.
type ICollectionLayoutAnchor interface {
	objectivec.IObject
	IsFractionalOffset() bool
	Edges() DirectionalRectEdge
	SetEdges(value IDirectionalRectEdge)
	IsAbsoluteOffset() bool
	SetIsAbsoluteOffset(value bool)
	Offset() coregraphics.CGPoint
	SetOffset(value coregraphics.CGPoint)
}

// An object that defines how to attach a supplementary item to an item in a collection view.
//
// You use an anchor to attach a supplementary item to a specific item. An anchor contains information about where on the item your supplementary item is attached, including: An edge or set of edges. You can attach a supplementary item to a single edge, or to a corner by specifying two adjacent edges. An offset from the item. By default, the supplementary item is anchored within the specified edges of the item it’s attached to. You can change this location by providing a custom offset when you create an anchor.


// An object that defines how to attach a supplementary item to an item in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor
type CollectionLayoutAnchor struct {
	objectivec.Object
}

// CollectionLayoutAnchorFrom constructs a [CollectionLayoutAnchor] from an unsafe.Pointer.
//
// An object that defines how to attach a supplementary item to an item in a collection view.
func CollectionLayoutAnchorFrom(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutAnchor) Autorelease() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutAnchor creates a new CollectionLayoutAnchor instance.
func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return getCollectionLayoutAnchorClass().New()
}



// A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isFractionalOffset
func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFractionalOffset"))
	return rv
}


// The edges of the item an anchor is attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/edges
func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := objc.Send[DirectionalRectEdge](c_.ID, objc.Sel("edges"))
	return rv
}


// The edges of the item an anchor is attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/edges
func (c_ CollectionLayoutAnchor) SetEdges(value IDirectionalRectEdge) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdges:"), value)
}


// A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/isabsoluteoffset
func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAbsoluteOffset"))
	return rv
}


// A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/isabsoluteoffset
func (c_ CollectionLayoutAnchor) SetIsAbsoluteOffset(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAbsoluteOffset:"), value)
}


// The floating-point value of the anchor’s offset from the item it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/offset
func (c_ CollectionLayoutAnchor) Offset() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("offset"))
	return rv
}


// The floating-point value of the anchor’s offset from the item it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutanchor/offset
func (c_ CollectionLayoutAnchor) SetOffset(value coregraphics.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOffset:"), value)
}



