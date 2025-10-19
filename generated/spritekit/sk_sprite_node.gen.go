// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKSpriteNode] class.
var (
	sKSpriteNodeClass     _SKSpriteNodeClass
	sKSpriteNodeClassOnce sync.Once
)

func getSKSpriteNodeClass() _SKSpriteNodeClass {
	sKSpriteNodeClassOnce.Do(func() {
		sKSpriteNodeClass = _SKSpriteNodeClass{objc.GetClass("SKSpriteNode")}
	})
	return sKSpriteNodeClass
}

type _SKSpriteNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKSpriteNode] class.
type ISKSpriteNode interface {
	ISKNode
}

// An image or solid color.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKSpriteNode
type SKSpriteNode struct {
	SKNode
}

// SKSpriteNodeFrom constructs a [SKSpriteNode] from an unsafe.Pointer.
//
// An image or solid color.
func SKSpriteNodeFrom(ptr unsafe.Pointer) SKSpriteNode {
	return SKSpriteNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKSpriteNodeClass) Alloc() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKSpriteNodeClass) New() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKSpriteNode) Init() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKSpriteNode) Autorelease() SKSpriteNode {
	rv := objc.Send[SKSpriteNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKSpriteNode creates a new SKSpriteNode instance.
func NewSKSpriteNode() SKSpriteNode {
	return getSKSpriteNodeClass().New()
}




