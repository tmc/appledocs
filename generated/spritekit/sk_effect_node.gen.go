// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKEffectNode] class.
var sKEffectNodeClass = _SKEffectNodeClass{objc.GetClass("SKEffectNode")}

type _SKEffectNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKEffectNode] class.
type ISKEffectNode interface {
	ISKNode
}

// A node that renders its children into a separate buffer, optionally applying an effect, before drawing the final result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKEffectNode

type SKEffectNode struct {
	SKNode
}

// SKEffectNodeFrom constructs a [SKEffectNode] from an unsafe.Pointer.
//
// A node that renders its children into a separate buffer, optionally applying an effect, before drawing the final result.
func SKEffectNodeFrom(ptr unsafe.Pointer) SKEffectNode {
	return SKEffectNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKEffectNodeClass) Alloc() SKEffectNode {
	rv := objc.Send[SKEffectNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKEffectNodeClass) New() SKEffectNode {
	rv := objc.Send[SKEffectNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKEffectNode) Init() SKEffectNode {
	rv := objc.Send[SKEffectNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKEffectNode) Autorelease() SKEffectNode {
	rv := objc.Send[SKEffectNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKEffectNode creates a new SKEffectNode instance.
func NewSKEffectNode() SKEffectNode {
	return sKEffectNodeClass.New()
}




