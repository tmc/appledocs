// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKCameraNode] class.
var sKCameraNodeClass = _SKCameraNodeClass{objc.GetClass("SKCameraNode")}

type _SKCameraNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKCameraNode] class.
type ISKCameraNode interface {
	ISKNode
}

// A node that determines which parts of the scene are visible within a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKCameraNode

type SKCameraNode struct {
	SKNode
}

// SKCameraNodeFrom constructs a [SKCameraNode] from an unsafe.Pointer.
//
// A node that determines which parts of the scene are visible within a view.
func SKCameraNodeFrom(ptr unsafe.Pointer) SKCameraNode {
	return SKCameraNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKCameraNodeClass) Alloc() SKCameraNode {
	rv := objc.Send[SKCameraNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKCameraNodeClass) New() SKCameraNode {
	rv := objc.Send[SKCameraNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKCameraNode) Init() SKCameraNode {
	rv := objc.Send[SKCameraNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKCameraNode) Autorelease() SKCameraNode {
	rv := objc.Send[SKCameraNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKCameraNode creates a new SKCameraNode instance.
func NewSKCameraNode() SKCameraNode {
	return sKCameraNodeClass.New()
}




