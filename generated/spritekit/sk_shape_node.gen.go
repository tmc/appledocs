// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKShapeNode] class.
var sKShapeNodeClass = _SKShapeNodeClass{objc.GetClass("SKShapeNode")}

type _SKShapeNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKShapeNode] class.
type ISKShapeNode interface {
	ISKNode
}

// A mathematical shape that can be stroked or filled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShapeNode

type SKShapeNode struct {
	SKNode
}

// SKShapeNodeFrom constructs a [SKShapeNode] from an unsafe.Pointer.
//
// A mathematical shape that can be stroked or filled.
func SKShapeNodeFrom(ptr unsafe.Pointer) SKShapeNode {
	return SKShapeNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKShapeNodeClass) Alloc() SKShapeNode {
	rv := objc.Send[SKShapeNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKShapeNodeClass) New() SKShapeNode {
	rv := objc.Send[SKShapeNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKShapeNode) Init() SKShapeNode {
	rv := objc.Send[SKShapeNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKShapeNode) Autorelease() SKShapeNode {
	rv := objc.Send[SKShapeNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKShapeNode creates a new SKShapeNode instance.
func NewSKShapeNode() SKShapeNode {
	return sKShapeNodeClass.New()
}




