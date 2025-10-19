// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKAction] class.
var sKActionClass = _SKActionClass{objc.GetClass("SKAction")}

type _SKActionClass struct {
	class objc.Class
}

// An interface definition for the [SKAction] class.
type ISKAction interface {
	objectivec.IObject
	ReversedAction() unsafe.Pointer
}

// An object that is run by a node to change its structure or content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAction

type SKAction struct {
	objectivec.Object
}

// SKActionFrom constructs a [SKAction] from an unsafe.Pointer.
//
// An object that is run by a node to change its structure or content.
func SKActionFrom(ptr unsafe.Pointer) SKAction {
	return SKAction{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKActionClass) Alloc() SKAction {
	rv := objc.Send[SKAction](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKActionClass) New() SKAction {
	rv := objc.Send[SKAction](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKAction) Init() SKAction {
	rv := objc.Send[SKAction](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKAction) Autorelease() SKAction {
	rv := objc.Send[SKAction](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAction creates a new SKAction instance.
func NewSKAction() SKAction {
	return sKActionClass.New()
}


// Creates an action that changes the alpha value of the node to . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAction/fadeOut(withDuration:)
func (sc _SKActionClass) FadeOutWithDuration(duration TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("fadeOutWithDuration:"), duration)
	return rv
}
// Creates an action that reverses the behavior of another action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKAction/reversed()
func (s_ SKAction) ReversedAction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("reversedAction"))
	return rv
}


