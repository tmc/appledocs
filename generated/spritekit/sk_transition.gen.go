// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTransition] class.
var (
	sKTransitionClass     _SKTransitionClass
	sKTransitionClassOnce sync.Once
)

func getSKTransitionClass() _SKTransitionClass {
	sKTransitionClassOnce.Do(func() {
		sKTransitionClass = _SKTransitionClass{objc.GetClass("SKTransition")}
	})
	return sKTransitionClass
}

type _SKTransitionClass struct {
	class objc.Class
}

// An interface definition for the [SKTransition] class.
type ISKTransition interface {
	objectivec.IObject
}

// An object used to perform an animated transition to a new scene. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTransition
type SKTransition struct {
	objectivec.Object
}

// SKTransitionFrom constructs a [SKTransition] from an unsafe.Pointer.
//
// An object used to perform an animated transition to a new scene.
func SKTransitionFrom(ptr unsafe.Pointer) SKTransition {
	return SKTransition{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTransitionClass) Alloc() SKTransition {
	rv := objc.Send[SKTransition](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTransitionClass) New() SKTransition {
	rv := objc.Send[SKTransition](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTransition) Init() SKTransition {
	rv := objc.Send[SKTransition](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTransition) Autorelease() SKTransition {
	rv := objc.Send[SKTransition](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTransition creates a new SKTransition instance.
func NewSKTransition() SKTransition {
	return getSKTransitionClass().New()
}


// Creates a transition that first fades to black and then fades to the new scene. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTransition/fade(withDuration:)
func (sc _SKTransitionClass) FadeWithDuration(sec TimeInterval) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("fadeWithDuration:"), sec)
	return rv
}


