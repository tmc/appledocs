// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlayerItemMediaDataCollector] class.
var (
	PlayerItemMediaDataCollectorClass     _PlayerItemMediaDataCollectorClass
	PlayerItemMediaDataCollectorClassOnce sync.Once
)

func getPlayerItemMediaDataCollectorClass() _PlayerItemMediaDataCollectorClass {
	PlayerItemMediaDataCollectorClassOnce.Do(func() {
		PlayerItemMediaDataCollectorClass = _PlayerItemMediaDataCollectorClass{objc.GetClass("AVPlayerItemMediaDataCollector")}
	})
	return PlayerItemMediaDataCollectorClass
}

type _PlayerItemMediaDataCollectorClass struct {
	class objc.Class
}





// An interface definition for the [PlayerItemMediaDataCollector] class.
type IPlayerItemMediaDataCollector interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemMediaDataCollectorClass) Alloc() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemMediaDataCollectorClass) New() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemMediaDataCollector) Init() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemMediaDataCollector) Autorelease() PlayerItemMediaDataCollector {
	rv := objc.Send[PlayerItemMediaDataCollector](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemMediaDataCollector creates a new PlayerItemMediaDataCollector instance.
func NewPlayerItemMediaDataCollector() PlayerItemMediaDataCollector {
	return getPlayerItemMediaDataCollectorClass().New()
}





// The abstract base for media data collectors.


// The abstract base for media data collectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemMediaDataCollector
type PlayerItemMediaDataCollector struct {
	objectivec.Object
}

// PlayerItemMediaDataCollectorFrom constructs a [PlayerItemMediaDataCollector] from an unsafe.Pointer.
//
// The abstract base for media data collectors.
func PlayerItemMediaDataCollectorFrom(ptr unsafe.Pointer) PlayerItemMediaDataCollector {
	return PlayerItemMediaDataCollector{objectivec.Object{objc.ID(ptr)}}
}































