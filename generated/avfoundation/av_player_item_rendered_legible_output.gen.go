// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerItemRenderedLegibleOutput] class.
var (
	PlayerItemRenderedLegibleOutputClass     _PlayerItemRenderedLegibleOutputClass
	PlayerItemRenderedLegibleOutputClassOnce sync.Once
)

func getPlayerItemRenderedLegibleOutputClass() _PlayerItemRenderedLegibleOutputClass {
	PlayerItemRenderedLegibleOutputClassOnce.Do(func() {
		PlayerItemRenderedLegibleOutputClass = _PlayerItemRenderedLegibleOutputClass{objc.GetClass("AVPlayerItemRenderedLegibleOutput")}
	})
	return PlayerItemRenderedLegibleOutputClass
}

type _PlayerItemRenderedLegibleOutputClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItemRenderedLegibleOutput] class.
type IPlayerItemRenderedLegibleOutput interface {
	objectivec.IObject
}

// A player item output that vends media with a legible characteristic as rendered pixel buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput
type PlayerItemRenderedLegibleOutput struct {
	objectivec.Object
}

// PlayerItemRenderedLegibleOutputFrom constructs a [PlayerItemRenderedLegibleOutput] from an unsafe.Pointer.
//
// A player item output that vends media with a legible characteristic as rendered pixel buffers.
func PlayerItemRenderedLegibleOutputFrom(ptr unsafe.Pointer) PlayerItemRenderedLegibleOutput {
	return PlayerItemRenderedLegibleOutput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemRenderedLegibleOutputClass) Alloc() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemRenderedLegibleOutputClass) New() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemRenderedLegibleOutput) Init() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemRenderedLegibleOutput) Autorelease() PlayerItemRenderedLegibleOutput {
	rv := objc.Send[PlayerItemRenderedLegibleOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemRenderedLegibleOutput creates a new PlayerItemRenderedLegibleOutput instance.
func NewPlayerItemRenderedLegibleOutput() PlayerItemRenderedLegibleOutput {
	return getPlayerItemRenderedLegibleOutputClass().New()
}




