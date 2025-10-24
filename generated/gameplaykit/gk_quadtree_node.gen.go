// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QuadtreeNode] class.
var (
	QuadtreeNodeClass     _QuadtreeNodeClass
	QuadtreeNodeClassOnce sync.Once
)

func getQuadtreeNodeClass() _QuadtreeNodeClass {
	QuadtreeNodeClassOnce.Do(func() {
		QuadtreeNodeClass = _QuadtreeNodeClass{objc.GetClass("GKQuadtreeNode")}
	})
	return QuadtreeNodeClass
}

type _QuadtreeNodeClass struct {
	class objc.Class
}

// An interface definition for the [QuadtreeNode] class.
type IQuadtreeNode interface {
	objectivec.IObject
	// properties:
	Quad() GKQuad /* not a class type */
	// methods:
}

// A helper class for managing the objects you organize in a quadtree.
//
// You don’t create instances of this class directly; instead, a object provides you with a instance when you add an element to a tree. If you plan to remove elements from the tree, keep references to the corresponding nodes so you can use the method for better performance. For more information, see .


// A helper class for managing the objects you organize in a quadtree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtreeNode
type QuadtreeNode struct {
	objectivec.Object
}

// QuadtreeNodeFrom constructs a [QuadtreeNode] from an unsafe.Pointer.
//
// A helper class for managing the objects you organize in a quadtree.
func QuadtreeNodeFrom(ptr unsafe.Pointer) QuadtreeNode {
	return QuadtreeNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QuadtreeNodeClass) Alloc() QuadtreeNode {
	rv := objc.Send[QuadtreeNode](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuadtreeNodeClass) New() QuadtreeNode {
	rv := objc.Send[QuadtreeNode](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuadtreeNode) Init() QuadtreeNode {
	rv := objc.Send[QuadtreeNode](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuadtreeNode) Autorelease() QuadtreeNode {
	rv := objc.Send[QuadtreeNode](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuadtreeNode creates a new QuadtreeNode instance.
func NewQuadtreeNode() QuadtreeNode {
	return getQuadtreeNodeClass().New()
}



// The axis-aligned bounding rectangle represented by the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKQuadtreeNode/quad
func (q_ QuadtreeNode) Quad() GKQuad /* not a class type */ {
	rv := objc.Send[Quad](q_.ID, objc.Sel("quad"))
	return rv
}



