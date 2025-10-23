// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerItemOutput] class.
var (
	PlayerItemOutputClass     _PlayerItemOutputClass
	PlayerItemOutputClassOnce sync.Once
)

func getPlayerItemOutputClass() _PlayerItemOutputClass {
	PlayerItemOutputClassOnce.Do(func() {
		PlayerItemOutputClass = _PlayerItemOutputClass{objc.GetClass("AVPlayerItemOutput")}
	})
	return PlayerItemOutputClass
}

type _PlayerItemOutputClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItemOutput] class.
type IPlayerItemOutput interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type PlayerItemOutput struct {
	objectivec.Object
}

// PlayerItemOutputFrom constructs a [PlayerItemOutput] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func PlayerItemOutputFrom(ptr unsafe.Pointer) PlayerItemOutput {
	return PlayerItemOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemOutputClass) Alloc() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemOutputClass) New() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemOutput) Init() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemOutput) Autorelease() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemOutput creates a new PlayerItemOutput instance.
func NewPlayerItemOutput() PlayerItemOutput {
	return getPlayerItemOutputClass().New()
}




