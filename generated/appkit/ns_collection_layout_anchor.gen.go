// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	Edges() DirectionalRectEdge
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
	Offset() corefoundation.CGPoint


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an anchor with the specified edges to attach to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:)
func NewCollectionLayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:"), edges)
	return rv
}


// Creates an anchor with the specified edges to attach to, offset by the provided absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:absoluteOffset:)
func NewCollectionLayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset corefoundation.CGPoint) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return rv
}


// Creates an anchor with the specified edges to attach to, offset by the provided fractional value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:fractionalOffset:)
func NewCollectionLayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset corefoundation.CGPoint) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return rv
}







// Creates an anchor with the specified edges to attach to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:"), edges)
	return rv
}


// Creates an anchor with the specified edges to attach to, offset by the provided absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:absoluteOffset:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return rv
}


// Creates an anchor with the specified edges to attach to, offset by the provided fractional value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:fractionalOffset:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return rv
}

















// The edges of the item an anchor is attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/edges
func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := objc.Send[DirectionalRectEdge](c_.ID, objc.Sel("edges"))
	return rv
}


// A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isAbsoluteOffset
func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAbsoluteOffset"))
	return rv
}


// A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isFractionalOffset
func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFractionalOffset"))
	return rv
}


// The floating-point value of the anchor’s offset from the item it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/offset
func (c_ CollectionLayoutAnchor) Offset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("offset"))
	return rv
}







