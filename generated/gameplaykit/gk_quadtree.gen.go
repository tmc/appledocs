// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Quadtree] class.
var (
	QuadtreeClass     _QuadtreeClass
	QuadtreeClassOnce sync.Once
)

func getQuadtreeClass() _QuadtreeClass {
	QuadtreeClassOnce.Do(func() {
		QuadtreeClass = _QuadtreeClass{objc.GetClass("GKQuadtree")}
	})
	return QuadtreeClass
}

type _QuadtreeClass struct {
	class objc.Class
}

// An interface definition for the [Quadtree] class.
type IQuadtree interface {
	objectivec.IObject
	AddElementWithPoint(element unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer
	AddElementWithQuad(element unsafe.Pointer, quad unsafe.Pointer) unsafe.Pointer
	ElementsAtPoint(point unsafe.Pointer) []appkit.NSObject
	ElementsInQuad(quad unsafe.Pointer) []appkit.NSObject
	RemoveElement(element unsafe.Pointer) bool
	RemoveElementWithNode(data unsafe.Pointer, node unsafe.Pointer) bool
}

// A data structure for organizing objects based on their locations in a two-dimensional space.
//
// A quadtree manages its structure to optimize for spatial searches—unlike a basic data structure such as an array or dictionary, a quadtree can find all elements occupying a specific position or region very quickly. The quadtree partitioning strategy divides space into four quadrants at each level, as illustrated in . When a quadrant contains more than one object, the tree subdivides that region into four smaller quadrants, adding a level to the tree. Quadtrees can be useful for many tasks in game design. For example: Deciding which game characters are close enough to each other for interaction Deciding which portions of a large game world need to be processed at a given time The class is one of three spatial partitioning data structures that GameplayKit provides. See these other classes for other tasks: The class provides the three-dimensional equivalent of a quadtree. Use an octree when you need to organize objects in 3D space. The class provides a different algorithm for two-dimensional spatial indexing. Quadtrees and R-trees have different performance tradeoffs for different tasks: quadtrees can be faster when objects are more uniformly distributed in space or when their positions change frequently, and R-trees can be faster when searching for all objects in a given region.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree
type Quadtree struct {
	objectivec.Object
}

// QuadtreeFrom constructs a [Quadtree] from an unsafe.Pointer.
//
// A data structure for organizing objects based on their locations in a two-dimensional space.
func QuadtreeFrom(ptr unsafe.Pointer) Quadtree {
	return Quadtree{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QuadtreeClass) Alloc() Quadtree {
	rv := objc.Send[Quadtree](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuadtreeClass) New() Quadtree {
	rv := objc.Send[Quadtree](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ Quadtree) Init() Quadtree {
	rv := objc.Send[Quadtree](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ Quadtree) Autorelease() Quadtree {
	rv := objc.Send[Quadtree](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuadtree creates a new Quadtree instance.
func NewQuadtree() Quadtree {
	return getQuadtreeClass().New()
}


// Initializes a quadtree with the specified dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/init(boundingQuad:minimumCellSize:)
func NewQuadtreeWithBoundingQuadMinimumCellSize(quad unsafe.Pointer, minCellSize unsafe.Pointer) Quadtree {
	instance := getQuadtreeClass().Alloc()
	rv := objc.Send[Quadtree](instance.ID, objc.Sel("initWithBoundingQuad:minimumCellSize:"), quad, minCellSize)
	rv.Autorelease()
	return rv
}


// Creates a quadtree with the specified dimensions.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/quadtreeWithBoundingQuad:minimumCellSize:
func (qc _QuadtreeClass) QuadtreeWithBoundingQuadMinimumCellSize(quad unsafe.Pointer, minCellSize unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(qc.class), objc.Sel("quadtreeWithBoundingQuad:minimumCellSize:"), quad, minCellSize)
	return rv
}

// Adds an object to the tree corresponding to the specified point in 2D space.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/add(_:at:)
func (q_ Quadtree) AddElementWithPoint(element unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("addElement:withPoint:"), element, point)
	return rv
}

// Adds an object to the tree corresponding to the specified region of 2D space.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/add(_:in:)
func (q_ Quadtree) AddElementWithQuad(element unsafe.Pointer, quad unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](q_.ID, objc.Sel("addElement:withQuad:"), element, quad)
	return rv
}

// Returns all objects whose corresponding locations overlap the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/elements(at:)
func (q_ Quadtree) ElementsAtPoint(point unsafe.Pointer) []appkit.NSObject {
	rv := objc.Send[[]appkit.NSObject](q_.ID, objc.Sel("elementsAtPoint:"), point)
	return rv
}

// Returns all objects whose corresponding locations overlap the specified region.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/elements(in:)
func (q_ Quadtree) ElementsInQuad(quad unsafe.Pointer) []appkit.NSObject {
	rv := objc.Send[[]appkit.NSObject](q_.ID, objc.Sel("elementsInQuad:"), quad)
	return rv
}

// Searches for the specified object and removes it from the tree.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/remove(_:)
func (q_ Quadtree) RemoveElement(element unsafe.Pointer) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("removeElement:"), element)
	return rv
}

// Removes the specified object from the tree, using a reference to its containing node.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/remove(_:using:)
func (q_ Quadtree) RemoveElementWithNode(data unsafe.Pointer, node unsafe.Pointer) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("removeElement:withNode:"), data, node)
	return rv
}


