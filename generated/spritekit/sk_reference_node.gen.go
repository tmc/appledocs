// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKReferenceNode] class.
var sKReferenceNodeClass = _SKReferenceNodeClass{objc.GetClass("SKReferenceNode")}

type _SKReferenceNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKReferenceNode] class.
type ISKReferenceNode interface {
	ISKNode
}

// A node that’s defined in an archived file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKReferenceNode

type SKReferenceNode struct {
	SKNode
}

// SKReferenceNodeFrom constructs a [SKReferenceNode] from an unsafe.Pointer.
//
// A node that’s defined in an archived file.
func SKReferenceNodeFrom(ptr unsafe.Pointer) SKReferenceNode {
	return SKReferenceNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKReferenceNodeClass) Alloc() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKReferenceNodeClass) New() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKReferenceNode) Init() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKReferenceNode) Autorelease() SKReferenceNode {
	rv := objc.Send[SKReferenceNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKReferenceNode creates a new SKReferenceNode instance.
func NewSKReferenceNode() SKReferenceNode {
	return sKReferenceNodeClass.New()
}




