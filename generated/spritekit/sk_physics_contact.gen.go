// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKPhysicsContact] class.
var (
	sKPhysicsContactClass     _SKPhysicsContactClass
	sKPhysicsContactClassOnce sync.Once
)

func getSKPhysicsContactClass() _SKPhysicsContactClass {
	sKPhysicsContactClassOnce.Do(func() {
		sKPhysicsContactClass = _SKPhysicsContactClass{objc.GetClass("SKPhysicsContact")}
	})
	return sKPhysicsContactClass
}

type _SKPhysicsContactClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsContact] class.
type ISKPhysicsContact interface {
	objectivec.IObject
}

// A description of the contact between two physics bodies.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsContact
type SKPhysicsContact struct {
	objectivec.Object
}

// SKPhysicsContactFrom constructs a [SKPhysicsContact] from an unsafe.Pointer.
//
// A description of the contact between two physics bodies.
func SKPhysicsContactFrom(ptr unsafe.Pointer) SKPhysicsContact {
	return SKPhysicsContact{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsContactClass) Alloc() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKPhysicsContactClass) New() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsContact) Init() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsContact) Autorelease() SKPhysicsContact {
	rv := objc.Send[SKPhysicsContact](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsContact creates a new SKPhysicsContact instance.
func NewSKPhysicsContact() SKPhysicsContact {
	return getSKPhysicsContactClass().New()
}




