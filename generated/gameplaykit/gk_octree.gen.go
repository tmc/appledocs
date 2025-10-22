// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Octree] class.
var (
	OctreeClass     _OctreeClass
	OctreeClassOnce sync.Once
)

func getOctreeClass() _OctreeClass {
	OctreeClassOnce.Do(func() {
		OctreeClass = _OctreeClass{objc.GetClass("GKOctree")}
	})
	return OctreeClass
}

type _OctreeClass struct {
	class objc.Class
}

// An interface definition for the [Octree] class.
type IOctree interface {
	objectivec.IObject
	AddElementWithPoint(element unsafe.Pointer, point unsafe.Pointer) OctreeNode
	AddElementWithBox(element unsafe.Pointer, box appkit.IBox) OctreeNode
	ElementsAtPoint(point unsafe.Pointer) []foundation.Object
	ElementsInBox(box appkit.IBox) []foundation.Object
	RemoveElement(element unsafe.Pointer) bool
	RemoveElementWithNode(element unsafe.Pointer, node IGKOctreeNode) bool
}

// A data structure for organizing objects based on their locations in a three-dimensional space.
//
// An octree manages its structure to optimize for spatial searches—unlike a basic data structure such as an array or dictionary, an octree can find all elements occupying a specific position or volume very quickly. The octree partitioning strategy divides space into eight octants at each level, as illustrated in . When an octant contains more than one object, the tree subdivides that region into eight smaller octants, adding a level to the tree. Octrees can be useful for many tasks in game design. For example: Deciding which game characters are close enough to each other for interaction Deciding which portions of a large game world need to be processed at a given time The class is one of three spatial partitioning data structures that GameplayKit provides, and the only one suited to three-dimensional data. See the class for the two-dimensional analogue of an octree, and the class for different ways to organize two-dimensional data.


// A data structure for organizing objects based on their locations in a three-dimensional space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree

type Octree struct {
	objectivec.Object
}

// OctreeFrom constructs a [Octree] from an unsafe.Pointer.
//
// A data structure for organizing objects based on their locations in a three-dimensional space.
func OctreeFrom(ptr unsafe.Pointer) Octree {
	return Octree{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OctreeClass) Alloc() Octree {
	rv := objc.Send[Octree](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OctreeClass) New() Octree {
	rv := objc.Send[Octree](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Octree) Init() Octree {
	rv := objc.Send[Octree](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Octree) Autorelease() Octree {
	rv := objc.Send[Octree](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOctree creates a new Octree instance.
func NewOctree() Octree {
	return getOctreeClass().New()
}




// Initializes an octree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/init(boundingBox:minimumCellSize:)

func NewOctreeWithBoundingBoxMinimumCellSize(box appkit.IBox, minCellSize float32) Octree {
	instance := getOctreeClass().Alloc()
	rv := objc.Send[Octree](instance.ID, objc.Sel("initWithBoundingBox:minimumCellSize:"), box, minCellSize)
	rv.Autorelease()
	return rv
}



// Creates an octree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/octreeWithBoundingBox:minimumCellSize:

func (oc _OctreeClass) OctreeWithBoundingBoxMinimumCellSize(box appkit.IBox, minCellSize float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("octreeWithBoundingBox:minimumCellSize:"), box, minCellSize)
	return rv
}



// Adds an object to the tree corresponding to the specified point in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/add(_:at:)

func (o_ Octree) AddElementWithPoint(element unsafe.Pointer, point unsafe.Pointer) OctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("addElement:withPoint:"), element, point)
	return rv
}



// Adds an object to the tree corresponding to the specified volume of 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/add(_:in:)

func (o_ Octree) AddElementWithBox(element unsafe.Pointer, box appkit.IBox) OctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("addElement:withBox:"), element, box)
	return rv
}



// Returns all objects whose corresponding locations overlap the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/elements(at:)

func (o_ Octree) ElementsAtPoint(point unsafe.Pointer) []foundation.Object {
	rv := objc.Send[[]foundation.Object](o_.ID, objc.Sel("elementsAtPoint:"), point)
	return rv
}



// Returns all objects whose corresponding locations overlap the specified volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/elements(in:)

func (o_ Octree) ElementsInBox(box appkit.IBox) []foundation.Object {
	rv := objc.Send[[]foundation.Object](o_.ID, objc.Sel("elementsInBox:"), box)
	return rv
}



// Searches for the specified object and removes it from the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/remove(_:)

func (o_ Octree) RemoveElement(element unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeElement:"), element)
	return rv
}



// Removes the specified object from the tree, using a reference to its containing node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/remove(_:using:)

func (o_ Octree) RemoveElementWithNode(element unsafe.Pointer, node IGKOctreeNode) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeElement:withNode:"), element, node)
	return rv
}


