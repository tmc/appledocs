// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKQuadtree */


/* debug [class_header]: Header for GKQuadtree */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Quadtree */
// An interface definition for the [Quadtree] class.
type IQuadtree interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Quadtree */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Quadtree */
	// methods:
	AddElementWithPoint(element objectivec.IObject, point objectivec.IObject) IQuadtreeNode
	AddElementWithQuad(element objectivec.IObject, quad objc.IObject /* cross-framework: GKQuad */) IQuadtreeNode
	ElementsAtPoint(point objectivec.IObject) []objc.IObject /* cross-framework: Object */
	ElementsInQuad(quad objc.IObject /* cross-framework: GKQuad */) []objc.IObject /* cross-framework: Object */
	RemoveElement(element objectivec.IObject) bool
	RemoveElementWithNode(data objectivec.IObject, node IGKQuadtreeNode) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Quadtree */
// Alloc allocates a new instance without initialization.
func (qc _QuadtreeClass) Alloc() Quadtree {
	rv := objc.Send[Quadtree](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Quadtree */
// A data structure for organizing objects based on their locations in a two-dimensional space.
//
// A quadtree manages its structure to optimize for spatial searches—unlike a basic data structure such as an array or dictionary, a quadtree can find all elements occupying a specific position or region very quickly. The quadtree partitioning strategy divides space into four quadrants at each level, as illustrated in . When a quadrant contains more than one object, the tree subdivides that region into four smaller quadrants, adding a level to the tree. Quadtrees can be useful for many tasks in game design. For example: Deciding which game characters are close enough to each other for interaction Deciding which portions of a large game world need to be processed at a given time The class is one of three spatial partitioning data structures that GameplayKit provides. See these other classes for other tasks: The class provides the three-dimensional equivalent of a quadtree. Use an octree when you need to organize objects in 3D space. The class provides a different algorithm for two-dimensional spatial indexing. Quadtrees and R-trees have different performance tradeoffs for different tasks: quadtrees can be faster when objects are more uniformly distributed in space or when their positions change frequently, and R-trees can be faster when searching for all objects in a given region.


// A data structure for organizing objects based on their locations in a two-dimensional space.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Quadtree */

// Initializes a quadtree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/init(boundingQuad:minimumCellSize:)
func NewQuadtreeWithBoundingQuadMinimumCellSize(quad objc.IObject /* cross-framework: GKQuad */, minCellSize float32) Quadtree {
	instance := getQuadtreeClass().Alloc()
	rv := objc.Send[Quadtree](instance.ID, objc.Sel("initWithBoundingQuad:minimumCellSize:"), quad, minCellSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewQuadtreeWithBoundingQuadMinimumCellSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Quadtree */

// Creates a quadtree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/quadtreeWithBoundingQuad:minimumCellSize:
func (qc _QuadtreeClass) QuadtreeWithBoundingQuadMinimumCellSize(quad objc.IObject /* cross-framework: GKQuad */, minCellSize float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(qc.class), objc.Sel("quadtreeWithBoundingQuad:minimumCellSize:"), quad, minCellSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=QuadtreeWithBoundingQuadMinimumCellSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Quadtree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Quadtree */

// Adds an object to the tree corresponding to the specified point in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/add(_:at:)
func (q_ Quadtree) AddElementWithPoint(element objectivec.IObject, point objectivec.IObject) IQuadtreeNode {
	rv := objc.Send[QuadtreeNode](q_.ID, objc.Sel("addElement:withPoint:"), element, point)
	return rv
}/* debug [instance_methods/method]: AddElementWithPoint */


// Adds an object to the tree corresponding to the specified region of 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/add(_:in:)
func (q_ Quadtree) AddElementWithQuad(element objectivec.IObject, quad objc.IObject /* cross-framework: GKQuad */) IQuadtreeNode {
	rv := objc.Send[QuadtreeNode](q_.ID, objc.Sel("addElement:withQuad:"), element, quad)
	return rv
}/* debug [instance_methods/method]: AddElementWithQuad */


// Returns all objects whose corresponding locations overlap the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/elements(at:)
func (q_ Quadtree) ElementsAtPoint(point objectivec.IObject) []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](q_.ID, objc.Sel("elementsAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: ElementsAtPoint */


// Returns all objects whose corresponding locations overlap the specified region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/elements(in:)
func (q_ Quadtree) ElementsInQuad(quad objc.IObject /* cross-framework: GKQuad */) []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](q_.ID, objc.Sel("elementsInQuad:"), quad)
	return rv
}/* debug [instance_methods/method]: ElementsInQuad */


// Searches for the specified object and removes it from the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/remove(_:)
func (q_ Quadtree) RemoveElement(element objectivec.IObject) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("removeElement:"), element)
	return rv
}/* debug [instance_methods/method]: RemoveElement */


// Removes the specified object from the tree, using a reference to its containing node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtree/remove(_:using:)
func (q_ Quadtree) RemoveElementWithNode(data objectivec.IObject, node IGKQuadtreeNode) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("removeElement:withNode:"), data, node)
	return rv
}/* debug [instance_methods/method]: RemoveElementWithNode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Quadtree */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKQuadtree */


