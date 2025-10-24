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
	SuppressesPlayerRendering() bool
	SetSuppressesPlayerRendering(value bool)


	

	// methods:
	ItemTimeForCVTimeStamp(timestamp objc.IObject /* cross-framework: TimeStamp */) objc.IObject /* cross-framework: Time */
	ItemTimeForHostTime(hostTimeInSeconds float64) objc.IObject /* cross-framework: Time */
	ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) objc.IObject /* cross-framework: Time */


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerItemOutputClass) Alloc() PlayerItemOutput {
	rv := objc.Send[PlayerItemOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An abstract class that defines the common interface to output media data from a player item.
//
// This class provides basic methods for converting time values to the timebase of the item. It also provides an option to suppress rendering of the output associated with the specific instance of this class.


// An abstract class that defines the common interface to output media data from a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput
type PlayerItemOutput struct {
	objectivec.Object
}

// PlayerItemOutputFrom constructs a [PlayerItemOutput] from an unsafe.Pointer.
//
// An abstract class that defines the common interface to output media data from a player item.
func PlayerItemOutputFrom(ptr unsafe.Pointer) PlayerItemOutput {
	return PlayerItemOutput{objectivec.Object{objc.ID(ptr)}}
}




















// Converts a Core Video timestamp to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(for:)
func (p_ PlayerItemOutput) ItemTimeForCVTimeStamp(timestamp objc.IObject /* cross-framework: TimeStamp */) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForCVTimeStamp:"), timestamp)
	return rv
}


// Converts a host time, specified in seconds, to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forHostTime:)
func (p_ PlayerItemOutput) ItemTimeForHostTime(hostTimeInSeconds float64) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForHostTime:"), hostTimeInSeconds)
	return rv
}


// Converts a Mach host time to the item’s timebase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forMachAbsoluteTime:)
func (p_ PlayerItemOutput) ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](p_.ID, objc.Sel("itemTimeForMachAbsoluteTime:"), machAbsoluteTime)
	return rv
}







// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/suppressesPlayerRendering
func (p_ PlayerItemOutput) SuppressesPlayerRendering() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("suppressesPlayerRendering"))
	return rv
}


// A Boolean value that indicates whether the player object renders the receiver’s output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/suppressesPlayerRendering
func (p_ PlayerItemOutput) SetSuppressesPlayerRendering(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSuppressesPlayerRendering:"), value)
}








