// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKRTree */


/* debug [class_header]: Header for GKRTree */
// The class instance for the [RTree] class.
var (
	RTreeClass     _RTreeClass
	RTreeClassOnce sync.Once
)

func getRTreeClass() _RTreeClass {
	RTreeClassOnce.Do(func() {
		RTreeClass = _RTreeClass{objc.GetClass("GKRTree")}
	})
	return RTreeClass
}

type _RTreeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RTree */
// An interface definition for the [RTree] class.
type IRTree interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RTree */
	// properties:
	QueryReserve() uint
	SetQueryReserve(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RTree */
	// methods:
	AddElementBoundingRectMinBoundingRectMaxSplitStrategy(element objectivec.IObject, boundingRectMin objectivec.IObject, boundingRectMax objectivec.IObject, splitStrategy RTreeSplitStrategy)
	ElementsInBoundingRectMinRectMax(rectMin objectivec.IObject, rectMax objectivec.IObject) []objc.IObject /* cross-framework: Object */
	RemoveElementBoundingRectMinBoundingRectMax(element objectivec.IObject, boundingRectMin objectivec.IObject, boundingRectMax objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RTree */
// Alloc allocates a new instance without initialization.
func (rc _RTreeClass) Alloc() RTree {
	rv := objc.Send[RTree](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RTreeClass) New() RTree {
	rv := objc.Send[RTree](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RTree) Init() RTree {
	rv := objc.Send[RTree](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RTree) Autorelease() RTree {
	rv := objc.Send[RTree](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRTree creates a new RTree instance.
func NewRTree() RTree {
	return getRTreeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RTree */
// A data structure that adaptively organizes objects based on their locations in a two-dimensional space.
//
// An R-tree manages its structure to optimize for spatial searches—unlike a basic data structure such as an array or dictionary, an R-tree can find all elements occupying a specific position or region very quickly. Additionally, R-trees adapt their internal structure as you add and remove elements, increasing the amount of time required to perform those operations, but decreasing the time required to search for elements later. An R-tree partitions the space it describes by calculating the minimum bounding regions that enclose each of the added objects. For example, in , the numbered shapes are objects added to the tree, and the rectangles marked with letters are the data structure the tree creates to organize them. In this example, the rectangle C is the smallest rectangle that entirely contains objects 1 and 2; the rectangle D is the smallest that contains objects 3, 4, and 5; the rectangle A is the smallest containing all the objects in rectangles C and D; and so on. The R-tree automatically creates these divisions in a way that keeps the tree balanced—that is, so that no branch of the tree contains significantly more objects or sub-branches than any other branch—so that searches for objects in the tree require a uniformly minimal amount of processing. R-trees can be useful for many tasks in game design. For example: Deciding which game characters are close enough to each other for interaction Deciding which portions of a large game world need to be processed at a given time Finding out which other objects are entirely contained within the region occupied by a certain object The class is one of three spatial partitioning data structures that GameplayKit provides. See these other classes for other tasks: The class provides the three-dimensional equivalent of a quadtree. Use an octree when you need to organize objects in 3D space. The class provides a different algorithm for two-dimensional spatial indexing. Quadtrees and R-trees have different performance tradeoffs for different tasks: quadtrees can be faster when objects are more uniformly distributed in space or when their positions change frequently, and R-trees can be faster when searching for all objects in a given region.


// A data structure that adaptively organizes objects based on their locations in a two-dimensional space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree
type RTree struct {
	objectivec.Object
}

// RTreeFrom constructs a [RTree] from an unsafe.Pointer.
//
// A data structure that adaptively organizes objects based on their locations in a two-dimensional space.
func RTreeFrom(ptr unsafe.Pointer) RTree {
	return RTree{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RTree */

// Initializes a new R-tree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/init(maxNumberOfChildren:)
func NewRTreeWithMaxNumberOfChildren(maxNumberOfChildren uint) RTree {
	instance := getRTreeClass().Alloc()
	rv := objc.Send[RTree](instance.ID, objc.Sel("initWithMaxNumberOfChildren:"), maxNumberOfChildren)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRTreeWithMaxNumberOfChildren */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RTree */

// Creates a new R-tree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/treeWithMaxNumberOfChildren:
func (rc _RTreeClass) TreeWithMaxNumberOfChildren(maxNumberOfChildren uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(rc.class), objc.Sel("treeWithMaxNumberOfChildren:"), maxNumberOfChildren)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TreeWithMaxNumberOfChildren) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RTree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RTree */

// Adds the specified object to the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/addElement(_:boundingRectMin:boundingRectMax:splitStrategy:)
func (r_ RTree) AddElementBoundingRectMinBoundingRectMaxSplitStrategy(element objectivec.IObject, boundingRectMin objectivec.IObject, boundingRectMax objectivec.IObject, splitStrategy RTreeSplitStrategy) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addElement:boundingRectMin:boundingRectMax:splitStrategy:"), element, boundingRectMin, boundingRectMax, splitStrategy)
}/* debug [instance_methods/method]: AddElementBoundingRectMinBoundingRectMaxSplitStrategy */


// Searches the tree and returns all elements found within the specified bounding region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/elements(inBoundingRectMin:rectMax:)
func (r_ RTree) ElementsInBoundingRectMinRectMax(rectMin objectivec.IObject, rectMax objectivec.IObject) []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](r_.ID, objc.Sel("elementsInBoundingRectMin:rectMax:"), rectMin, rectMax)
	return rv
}/* debug [instance_methods/method]: ElementsInBoundingRectMinRectMax */


// Removes the specified object from the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/removeElement(_:boundingRectMin:boundingRectMax:)
func (r_ RTree) RemoveElementBoundingRectMinBoundingRectMax(element objectivec.IObject, boundingRectMin objectivec.IObject, boundingRectMax objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeElement:boundingRectMin:boundingRectMax:"), element, boundingRectMin, boundingRectMax)
}/* debug [instance_methods/method]: RemoveElementBoundingRectMinBoundingRectMax */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RTree */

// The number of elements to reserve space for when searching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/queryReserve
func (r_ RTree) QueryReserve() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("queryReserve"))
	return rv
}/* debug [instance_properties/getter]: queryReserve */


// The number of elements to reserve space for when searching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRTree/queryReserve
func (r_ RTree) SetQueryReserve(value uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setQueryReserve:"), value)
}/* debug [instance_properties/setter]: queryReserve */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKRTree */


