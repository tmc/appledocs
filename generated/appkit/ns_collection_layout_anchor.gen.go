// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutAnchor */


/* debug [class_header]: Header for NSCollectionLayoutAnchor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutAnchor */
// An interface definition for the [CollectionLayoutAnchor] class.
type ICollectionLayoutAnchor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutAnchor */
	// properties:
	Edges() DirectionalRectEdge
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
	Offset() vision.Point
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutAnchor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutAnchor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutAnchor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutAnchor */

// Creates an anchor with the specified edges to attach to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:)
func NewCollectionLayoutAnchorWithEdges(edges DirectionalRectEdge) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:"), edges)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutAnchorWithEdges */


// Creates an anchor with the specified edges to attach to, offset by the provided absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:absoluteOffset:)
func NewCollectionLayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset vision.Point) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutAnchorWithEdgesAbsoluteOffset */


// Creates an anchor with the specified edges to attach to, offset by the provided fractional value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:fractionalOffset:)
func NewCollectionLayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset vision.Point) CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(getCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutAnchorWithEdgesFractionalOffset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutAnchor */

// Creates an anchor with the specified edges to attach to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdges(edges DirectionalRectEdge) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:"), edges)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAnchorWithEdges) */


// Creates an anchor with the specified edges to attach to, offset by the provided absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:absoluteOffset:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesAbsoluteOffset(edges DirectionalRectEdge, absoluteOffset vision.Point) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAnchorWithEdgesAbsoluteOffset) */


// Creates an anchor with the specified edges to attach to, offset by the provided fractional value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:fractionalOffset:)
func (cc _CollectionLayoutAnchorClass) LayoutAnchorWithEdgesFractionalOffset(edges DirectionalRectEdge, fractionalOffset vision.Point) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayoutAnchorWithEdgesFractionalOffset) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutAnchor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutAnchor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutAnchor */

// The edges of the item an anchor is attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/edges
func (c_ CollectionLayoutAnchor) Edges() DirectionalRectEdge {
	rv := objc.Send[DirectionalRectEdge](c_.ID, objc.Sel("edges"))
	return rv
}/* debug [instance_properties/getter]: edges */


// A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isAbsoluteOffset
func (c_ CollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAbsoluteOffset"))
	return rv
}/* debug [instance_properties/getter]: isAbsoluteOffset */


// A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isFractionalOffset
func (c_ CollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFractionalOffset"))
	return rv
}/* debug [instance_properties/getter]: isFractionalOffset */


// The floating-point value of the anchor’s offset from the item it’s attached to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/offset
func (c_ CollectionLayoutAnchor) Offset() vision.Point {
	rv := objc.Send[vision.Point](c_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutAnchor */


