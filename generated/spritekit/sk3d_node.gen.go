// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SK3DNode] class.
var sK3DNodeClass = _SK3DNodeClass{objc.GetClass("SK3DNode")}

type _SK3DNodeClass struct {
	class objc.Class
}

// An interface definition for the [SK3DNode] class.
type ISK3DNode interface {
	ISKNode
}

// 3D SceneKit content drawn as a flattened sprite. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SK3DNode

type SK3DNode struct {
	SKNode
}

// SK3DNodeFrom constructs a [SK3DNode] from an unsafe.Pointer.
//
// 3D SceneKit content drawn as a flattened sprite.
func SK3DNodeFrom(ptr unsafe.Pointer) SK3DNode {
	return SK3DNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SK3DNodeClass) Alloc() SK3DNode {
	rv := objc.Send[SK3DNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return sK3DNodeClass.New()
}




