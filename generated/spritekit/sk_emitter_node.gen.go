// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKEmitterNode] class.
var (
	sKEmitterNodeClass     _SKEmitterNodeClass
	sKEmitterNodeClassOnce sync.Once
)

func getSKEmitterNodeClass() _SKEmitterNodeClass {
	sKEmitterNodeClassOnce.Do(func() {
		sKEmitterNodeClass = _SKEmitterNodeClass{objc.GetClass("SKEmitterNode")}
	})
	return sKEmitterNodeClass
}

type _SKEmitterNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKEmitterNode] class.
type ISKEmitterNode interface {
	ISKNode
}

// A source of various particle effects.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode
type SKEmitterNode struct {
	SKNode
}

// SKEmitterNodeFrom constructs a [SKEmitterNode] from an unsafe.Pointer.
//
// A source of various particle effects.
func SKEmitterNodeFrom(ptr unsafe.Pointer) SKEmitterNode {
	return SKEmitterNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKEmitterNodeClass) Alloc() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKEmitterNodeClass) New() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKEmitterNode) Init() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKEmitterNode) Autorelease() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKEmitterNode creates a new SKEmitterNode instance.
func NewSKEmitterNode() SKEmitterNode {
	return getSKEmitterNodeClass().New()
}




