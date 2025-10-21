// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [OctreeNode] class.
type IOctreeNode interface {
	objectivec.IObject
}

// A helper class for managing the objects you organize in an octree.
//
// You don’t create instances of this class directly; instead, a object provides you with a instance when you add an element to a tree. If you plan to remove elements from the tree, keep references to the corresponding nodes so you can use the method for better performance.
//
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

// Alloc allocates a new instance without initialization.
func (oc _OctreeNodeClass) Alloc() OctreeNode {
	rv := objc.Send[OctreeNode](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The axis-aligned bounding box represented by the node.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKOctreeNode/box
func (o_ OctreeNode) Box() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("box"))
	return rv
}



