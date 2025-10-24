// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKOctree */


/* debug [class_header]: Header for GKOctree */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Octree */
// An interface definition for the [Octree] class.
type IOctree interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Octree */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Octree */
	// methods:
	AddElementWithPoint(element objectivec.IObject, point objectivec.IObject) IOctreeNode
	AddElementWithBox(element objectivec.IObject, box objc.IObject /* cross-framework: GKBox */) IOctreeNode
	ElementsAtPoint(point objectivec.IObject) []objc.IObject /* cross-framework: Object */
	ElementsInBox(box objc.IObject /* cross-framework: GKBox */) []objc.IObject /* cross-framework: Object */
	RemoveElement(element objectivec.IObject) bool
	RemoveElementWithNode(element objectivec.IObject, node IGKOctreeNode) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Octree */
// Alloc allocates a new instance without initialization.
func (oc _OctreeClass) Alloc() Octree {
	rv := objc.Send[Octree](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Octree */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Octree */

// Initializes an octree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/init(boundingBox:minimumCellSize:)
func NewOctreeWithBoundingBoxMinimumCellSize(box objc.IObject /* cross-framework: GKBox */, minCellSize float32) Octree {
	instance := getOctreeClass().Alloc()
	rv := objc.Send[Octree](instance.ID, objc.Sel("initWithBoundingBox:minimumCellSize:"), box, minCellSize)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOctreeWithBoundingBoxMinimumCellSize */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Octree */

// Creates an octree with the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/octreeWithBoundingBox:minimumCellSize:
func (oc _OctreeClass) OctreeWithBoundingBoxMinimumCellSize(box objc.IObject /* cross-framework: GKBox */, minCellSize float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(oc.class), objc.Sel("octreeWithBoundingBox:minimumCellSize:"), box, minCellSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OctreeWithBoundingBoxMinimumCellSize) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Octree */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Octree */

// Adds an object to the tree corresponding to the specified point in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/add(_:at:)
func (o_ Octree) AddElementWithPoint(element objectivec.IObject, point objectivec.IObject) IOctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("addElement:withPoint:"), element, point)
	return rv
}/* debug [instance_methods/method]: AddElementWithPoint */


// Adds an object to the tree corresponding to the specified volume of 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/add(_:in:)
func (o_ Octree) AddElementWithBox(element objectivec.IObject, box objc.IObject /* cross-framework: GKBox */) IOctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("addElement:withBox:"), element, box)
	return rv
}/* debug [instance_methods/method]: AddElementWithBox */


// Returns all objects whose corresponding locations overlap the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/elements(at:)
func (o_ Octree) ElementsAtPoint(point objectivec.IObject) []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](o_.ID, objc.Sel("elementsAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: ElementsAtPoint */


// Returns all objects whose corresponding locations overlap the specified volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/elements(in:)
func (o_ Octree) ElementsInBox(box objc.IObject /* cross-framework: GKBox */) []objc.IObject /* cross-framework: Object */ {
	rv := objc.Send[[]foundation.Object](o_.ID, objc.Sel("elementsInBox:"), box)
	return rv
}/* debug [instance_methods/method]: ElementsInBox */


// Searches for the specified object and removes it from the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/remove(_:)
func (o_ Octree) RemoveElement(element objectivec.IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeElement:"), element)
	return rv
}/* debug [instance_methods/method]: RemoveElement */


// Removes the specified object from the tree, using a reference to its containing node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctree/remove(_:using:)
func (o_ Octree) RemoveElementWithNode(element objectivec.IObject, node IGKOctreeNode) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("removeElement:withNode:"), element, node)
	return rv
}/* debug [instance_methods/method]: RemoveElementWithNode */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Octree */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKOctree */


