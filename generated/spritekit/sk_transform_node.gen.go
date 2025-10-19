// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKTransformNode] class.
var (
	sKTransformNodeClass     _SKTransformNodeClass
	sKTransformNodeClassOnce sync.Once
)

func getSKTransformNodeClass() _SKTransformNodeClass {
	sKTransformNodeClassOnce.Do(func() {
		sKTransformNodeClass = _SKTransformNodeClass{objc.GetClass("SKTransformNode")}
	})
	return sKTransformNodeClass
}

type _SKTransformNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKTransformNode] class.
type ISKTransformNode interface {
	ISKNode
}

// A node that allows its children to rotate in 3D.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTransformNode
type SKTransformNode struct {
	SKNode
}

// SKTransformNodeFrom constructs a [SKTransformNode] from an unsafe.Pointer.
//
// A node that allows its children to rotate in 3D.
func SKTransformNodeFrom(ptr unsafe.Pointer) SKTransformNode {
	return SKTransformNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTransformNodeClass) Alloc() SKTransformNode {
	rv := objc.Send[SKTransformNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTransformNodeClass) New() SKTransformNode {
	rv := objc.Send[SKTransformNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTransformNode) Init() SKTransformNode {
	rv := objc.Send[SKTransformNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTransformNode) Autorelease() SKTransformNode {
	rv := objc.Send[SKTransformNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTransformNode creates a new SKTransformNode instance.
func NewSKTransformNode() SKTransformNode {
	return getSKTransformNodeClass().New()
}




