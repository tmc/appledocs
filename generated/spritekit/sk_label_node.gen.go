// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKLabelNode] class.
var (
	sKLabelNodeClass     _SKLabelNodeClass
	sKLabelNodeClassOnce sync.Once
)

func getSKLabelNodeClass() _SKLabelNodeClass {
	sKLabelNodeClassOnce.Do(func() {
		sKLabelNodeClass = _SKLabelNodeClass{objc.GetClass("SKLabelNode")}
	})
	return sKLabelNodeClass
}

type _SKLabelNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKLabelNode] class.
type ISKLabelNode interface {
	ISKNode
}

// A graphical element that draws text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKLabelNode
type SKLabelNode struct {
	SKNode
}

// SKLabelNodeFrom constructs a [SKLabelNode] from an unsafe.Pointer.
//
// A graphical element that draws text.
func SKLabelNodeFrom(ptr unsafe.Pointer) SKLabelNode {
	return SKLabelNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKLabelNodeClass) Alloc() SKLabelNode {
	rv := objc.Send[SKLabelNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKLabelNodeClass) New() SKLabelNode {
	rv := objc.Send[SKLabelNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKLabelNode) Init() SKLabelNode {
	rv := objc.Send[SKLabelNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKLabelNode) Autorelease() SKLabelNode {
	rv := objc.Send[SKLabelNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKLabelNode creates a new SKLabelNode instance.
func NewSKLabelNode() SKLabelNode {
	return getSKLabelNodeClass().New()
}




