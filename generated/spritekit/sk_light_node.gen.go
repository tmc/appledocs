// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKLightNode] class.
var (
	sKLightNodeClass     _SKLightNodeClass
	sKLightNodeClassOnce sync.Once
)

func getSKLightNodeClass() _SKLightNodeClass {
	sKLightNodeClassOnce.Do(func() {
		sKLightNodeClass = _SKLightNodeClass{objc.GetClass("SKLightNode")}
	})
	return sKLightNodeClass
}

type _SKLightNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKLightNode] class.
type ISKLightNode interface {
	ISKNode
}

// A node that lights surrounding nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKLightNode
type SKLightNode struct {
	SKNode
}

// SKLightNodeFrom constructs a [SKLightNode] from an unsafe.Pointer.
//
// A node that lights surrounding nodes.
func SKLightNodeFrom(ptr unsafe.Pointer) SKLightNode {
	return SKLightNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKLightNodeClass) Alloc() SKLightNode {
	rv := objc.Send[SKLightNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKLightNodeClass) New() SKLightNode {
	rv := objc.Send[SKLightNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKLightNode) Init() SKLightNode {
	rv := objc.Send[SKLightNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKLightNode) Autorelease() SKLightNode {
	rv := objc.Send[SKLightNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKLightNode creates a new SKLightNode instance.
func NewSKLightNode() SKLightNode {
	return getSKLightNodeClass().New()
}




