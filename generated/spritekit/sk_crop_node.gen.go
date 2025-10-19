// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKCropNode] class.
var (
	sKCropNodeClass     _SKCropNodeClass
	sKCropNodeClassOnce sync.Once
)

func getSKCropNodeClass() _SKCropNodeClass {
	sKCropNodeClassOnce.Do(func() {
		sKCropNodeClass = _SKCropNodeClass{objc.GetClass("SKCropNode")}
	})
	return sKCropNodeClass
}

type _SKCropNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKCropNode] class.
type ISKCropNode interface {
	ISKNode
}

// A node that masks pixels drawn by its children so that only some pixels are seen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKCropNode
type SKCropNode struct {
	SKNode
}

// SKCropNodeFrom constructs a [SKCropNode] from an unsafe.Pointer.
//
// A node that masks pixels drawn by its children so that only some pixels are seen.
func SKCropNodeFrom(ptr unsafe.Pointer) SKCropNode {
	return SKCropNode{
		SKNode: SKNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKCropNodeClass) Alloc() SKCropNode {
	rv := objc.Send[SKCropNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKCropNodeClass) New() SKCropNode {
	rv := objc.Send[SKCropNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKCropNode) Init() SKCropNode {
	rv := objc.Send[SKCropNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKCropNode) Autorelease() SKCropNode {
	rv := objc.Send[SKCropNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKCropNode creates a new SKCropNode instance.
func NewSKCropNode() SKCropNode {
	return getSKCropNodeClass().New()
}




