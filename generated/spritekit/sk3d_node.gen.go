// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SK3DNode] class.
var (
	sK3DNodeClass     _SK3DNodeClass
	sK3DNodeClassOnce sync.Once
)

func getSK3DNodeClass() _SK3DNodeClass {
	sK3DNodeClassOnce.Do(func() {
		sK3DNodeClass = _SK3DNodeClass{objc.GetClass("SK3DNode")}
	})
	return sK3DNodeClass
}

type _SK3DNodeClass struct {
	class objc.Class
}

// An interface definition for the [SK3DNode] class.
type ISK3DNode interface {
	INode
}

// 3D SceneKit content drawn as a flattened sprite.
//
// Use objects to incorporate 3D SceneKit content into a SpriteKit-based game. When SpriteKit renders the node, the SceneKit scene is animated and rendered first. Then this rendered image is composited into the SpriteKit scene. Use the property to specify the SceneKit scene to be rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SK3DNode
type SK3DNode struct {
	Node
}

// SK3DNodeFrom constructs a [SK3DNode] from an unsafe.Pointer.
//
// 3D SceneKit content drawn as a flattened sprite.
func SK3DNodeFrom(ptr unsafe.Pointer) SK3DNode {
	return SK3DNode{
		Node: NodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SK3DNodeClass) Alloc() SK3DNode {
	rv := objc.Send[SK3DNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SK3DNodeClass) New() SK3DNode {
	rv := objc.Send[SK3DNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SK3DNode) Init() SK3DNode {
	rv := objc.Send[SK3DNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SK3DNode) Autorelease() SK3DNode {
	rv := objc.Send[SK3DNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSK3DNode creates a new SK3DNode instance.
func NewSK3DNode() SK3DNode {
	return getSK3DNodeClass().New()
}




