// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKOctreeNode */


/* debug [class_header]: Header for GKOctreeNode */
// The class instance for the [OctreeNode] class.
var (
	OctreeNodeClass     _OctreeNodeClass
	OctreeNodeClassOnce sync.Once
)

func getOctreeNodeClass() _OctreeNodeClass {
	OctreeNodeClassOnce.Do(func() {
		OctreeNodeClass = _OctreeNodeClass{objc.GetClass("GKOctreeNode")}
	})
	return OctreeNodeClass
}

type _OctreeNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OctreeNode */
// An interface definition for the [OctreeNode] class.
type IOctreeNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OctreeNode */
	// properties:
	Box() objc.IObject /* cross-framework: GKBox */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OctreeNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OctreeNode */
// Alloc allocates a new instance without initialization.
func (oc _OctreeNodeClass) Alloc() OctreeNode {
	rv := objc.Send[OctreeNode](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OctreeNodeClass) New() OctreeNode {
	rv := objc.Send[OctreeNode](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OctreeNode) Init() OctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OctreeNode) Autorelease() OctreeNode {
	rv := objc.Send[OctreeNode](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOctreeNode creates a new OctreeNode instance.
func NewOctreeNode() OctreeNode {
	return getOctreeNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OctreeNode */
// A helper class for managing the objects you organize in an octree.
//
// You don’t create instances of this class directly; instead, a object provides you with a instance when you add an element to a tree. If you plan to remove elements from the tree, keep references to the corresponding nodes so you can use the method for better performance.


// A helper class for managing the objects you organize in an octree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctreeNode
type OctreeNode struct {
	objectivec.Object
}

// OctreeNodeFrom constructs a [OctreeNode] from an unsafe.Pointer.
//
// A helper class for managing the objects you organize in an octree.
func OctreeNodeFrom(ptr unsafe.Pointer) OctreeNode {
	return OctreeNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OctreeNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OctreeNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OctreeNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OctreeNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OctreeNode */

// The axis-aligned bounding box represented by the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctreeNode/box
func (o_ OctreeNode) Box() objc.IObject /* cross-framework: GKBox */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("box"))
	return rv
}/* debug [instance_properties/getter]: box */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKOctreeNode */



