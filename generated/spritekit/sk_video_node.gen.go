// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKVideoNode] class.
var (
	sKVideoNodeClass     _SKVideoNodeClass
	sKVideoNodeClassOnce sync.Once
)

func getSKVideoNodeClass() _SKVideoNodeClass {
	sKVideoNodeClassOnce.Do(func() {
		sKVideoNodeClass = _SKVideoNodeClass{objc.GetClass("SKVideoNode")}
	})
	return sKVideoNodeClass
}

type _SKVideoNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKVideoNode] class.
type ISKVideoNode interface {
	ISKNode
}

// A graphical element that plays video content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKVideoNode
type SKVideoNode struct {
	SKNode
}

// SKVideoNodeFrom constructs a [SKVideoNode] from an unsafe.Pointer.
//
// A graphical element that plays video content.
func SKVideoNodeFrom(ptr unsafe.Pointer) SKVideoNode {
	return SKVideoNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKVideoNodeClass) Alloc() SKVideoNode {
	rv := objc.Send[SKVideoNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKVideoNodeClass) New() SKVideoNode {
	rv := objc.Send[SKVideoNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKVideoNode) Init() SKVideoNode {
	rv := objc.Send[SKVideoNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKVideoNode) Autorelease() SKVideoNode {
	rv := objc.Send[SKVideoNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKVideoNode creates a new SKVideoNode instance.
func NewSKVideoNode() SKVideoNode {
	return getSKVideoNodeClass().New()
}




