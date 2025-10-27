// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CoordinatedPlaybackSuspension] class.
var (
	CoordinatedPlaybackSuspensionClass     _CoordinatedPlaybackSuspensionClass
	CoordinatedPlaybackSuspensionClassOnce sync.Once
)

func getCoordinatedPlaybackSuspensionClass() _CoordinatedPlaybackSuspensionClass {
	CoordinatedPlaybackSuspensionClassOnce.Do(func() {
		CoordinatedPlaybackSuspensionClass = _CoordinatedPlaybackSuspensionClass{objc.GetClass("AVCoordinatedPlaybackSuspension")}
	})
	return CoordinatedPlaybackSuspensionClass
}

type _CoordinatedPlaybackSuspensionClass struct {
	class objc.Class
}





// An interface definition for the [CoordinatedPlaybackSuspension] class.
type ICoordinatedPlaybackSuspension interface {
	objectivec.IObject
	

	// properties:
	BeginDate() foundation.foundation.INSDate
	Reason() CoordinatedPlaybackSuspensionReason


	

	// methods:
	End()
	EndProposingNewTime(time objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (cc _CoordinatedPlaybackSuspensionClass) Alloc() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CoordinatedPlaybackSuspensionClass) New() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoordinatedPlaybackSuspension) Init() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoordinatedPlaybackSuspension) Autorelease() CoordinatedPlaybackSuspension {
	rv := objc.Send[CoordinatedPlaybackSuspension](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoordinatedPlaybackSuspension creates a new CoordinatedPlaybackSuspension instance.
func NewCoordinatedPlaybackSuspension() CoordinatedPlaybackSuspension {
	return getCoordinatedPlaybackSuspensionClass().New()
}





// An object that represents a temporary suspension of coordinated playback.
//
// See the playback coordinator’s method for details about suspending playback.


// An object that represents a temporary suspension of coordinated playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension
type CoordinatedPlaybackSuspension struct {
	objectivec.Object
}

// CoordinatedPlaybackSuspensionFrom constructs a [CoordinatedPlaybackSuspension] from an unsafe.Pointer.
//
// An object that represents a temporary suspension of coordinated playback.
func CoordinatedPlaybackSuspensionFrom(ptr unsafe.Pointer) CoordinatedPlaybackSuspension {
	return CoordinatedPlaybackSuspension{objectivec.Object{objc.ID(ptr)}}
}




















// Ends a suspension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end()
func (c_ CoordinatedPlaybackSuspension) End() {
	objc.Send[objc.ID](c_.ID, objc.Sel("end"))
}


// Ends a suspension and proposes a new playback time to the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end(proposingNewTime:)
func (c_ CoordinatedPlaybackSuspension) EndProposingNewTime(time objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("endProposingNewTime:"), time)
}







// The time the suspension begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/beginDate
func (c_ CoordinatedPlaybackSuspension) BeginDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("beginDate"))
	return rv
}


// The reason for the suspension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/reason-swift.property
func (c_ CoordinatedPlaybackSuspension) Reason() CoordinatedPlaybackSuspensionReason {
	rv := objc.Send[CoordinatedPlaybackSuspensionReason](c_.ID, objc.Sel("reason"))
	return rv
}








