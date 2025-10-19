// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKAudioNode] class.
var sKAudioNodeClass = _SKAudioNodeClass{objc.GetClass("SKAudioNode")}

type _SKAudioNodeClass struct {
	class objc.Class
}

// An interface definition for the [SKAudioNode] class.
type ISKAudioNode interface {
	ISKNode
}

// A node that plays audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAudioNode

type SKAudioNode struct {
	SKNode
}

// SKAudioNodeFrom constructs a [SKAudioNode] from an unsafe.Pointer.
//
// A node that plays audio.
func SKAudioNodeFrom(ptr unsafe.Pointer) SKAudioNode {
	return SKAudioNode{
		SKNode: SKNodeFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKAudioNodeClass) Alloc() SKAudioNode {
	rv := objc.Send[SKAudioNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKAudioNodeClass) New() SKAudioNode {
	rv := objc.Send[SKAudioNode](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKAudioNode) Init() SKAudioNode {
	rv := objc.Send[SKAudioNode](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKAudioNode) Autorelease() SKAudioNode {
	rv := objc.Send[SKAudioNode](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAudioNode creates a new SKAudioNode instance.
func NewSKAudioNode() SKAudioNode {
	return sKAudioNodeClass.New()
}




